package macro

import "ch8/io"

// defcons
const (
	MAXCHARS int = 5000     // size of name-defn table
	MAXDEF   int = MAXCHARS // max chars in a defn
	MAXTOK   int = MAXCHARS // max chars in a token
)

const BLANK uint8 = ' '
const COMMA uint8 = ','
const LPAREN uint8 = '('
const RPAREN uint8 = ')'
const ENDSTR uint8 = '\n'

const (
	defname  = "define"
	null     = ""
	exprname = "expr"
	subname  = "substr"
	ifname   = "ifelse"
	lenname  = "len"
	chqname  = "changeq"
)

var lquote = '`'
var rquote = '\''

// macro -- expand macros with arguments
func Macro() {
	initmacro()
	install(defname, null, DEFTYPE)
	install(exprname, null, EXPRTYPE)
	install(subname, null, SUBTYPE)
	install(ifname, null, IFTYPE)
	install(lenname, null, LENTYPE)
	install(chqname, null, CHQTYPE)

	cp = 0
	ap = 1
	ep = 1

	var defn string
	var toktype sttype // type returned by lookup
	for token := gettok(MAXTOK); len(token) > 0; token = gettok(MAXTOK) {
		if isletter(token[0]) {
			if !lookup(token, &defn, &toktype) {
				puttok(token)
			} else {
				cp++
				if cp > CALLSIZE {
					io.Error("macro: call stack overflow")
				}
				callstk[cp] = ap
				typestk[cp] = toktype
				ap = push(ep, &argstk, ap)
				puttok(defn)   // push definition
				putchr(ENDSTR) // CHECK: do we need it?

				ap = push(ep, &argstk, ap)
				puttok(token)  // stack name
				putchr(ENDSTR) // CHECK: do we need it?

				ap = push(ep, &argstk, ap)
				token = gettok(MAXTOK) // peek at next
				t := token[0]
				pbstr(token)
				if t != LPAREN { // add ()
					putback(RPAREN)
					putback(LPAREN)
				}
				plev[cp] = 0
			}
		} else if token[0] == lquote { // strip quotes
			nplar := 1
			for nplar != 0 {
				token = gettok(MAXTOK)
				if len(token) == 0 {
					io.Error("macro: missing right quote")
				}
				t := token[0]
				if t == rquote {
					nplar--
				} else if t == lquote {
					nplar++
				}
				if nplar > 0 {
					puttok(token)
				}
			}
		} else if cp == 0 { // not in a macro at all
			puttok(token)
		} else if token[0] == LPAREN {
			if plev[cp] > 0 {
				puttok(token)
			}
			plev[cp]++
		} else if token[0] == RPAREN {
			plev[cp]--
			if plev[cp] > 0 {
				puttok(token)
			} else { // end of argument list
				putchr(ENDSTR)
				eval(&argstk, typestk[cp], callstk[cp], ap-1)
				ap = callstk[cp] // pop eval stack
				ep = argstk[ap]
				cp--
			}

		} else if token[0] == COMMA && plev[cp] == 1 {
			putchr(ENDSTR) //new argument
			ap = push(ep, &argstk, ap)
		} else {
			puttok(token) // just stack it
		}
	}
	if cp != 0 {
		io.Error("macro: unexpected end of input")
	}
}

// initmacro -- initialize variables for macro
func initmacro() {
	bp = 0 // pushback buffer pointer
	inithash()
	// lquote = ord(GRAVE)
	// rquote = ord(ACUTE)
}

// puttok -- put token on output or evaluation stack
func puttok(s string) {
	for _, ch := range s {
		// we expect only ASCII
		putchr(byte(ch))
	}
}

// putchr -- put single char on output or evaluation stack
func putchr(c byte) {
	if cp <= 0 {
		io.Putc(c)
	} else {
		if ep > EVALSIZE {
			io.Error("macro: evaluation stack overflow")
		}
		evalstk[ep] = c
		ep++
	}
}

// push -- push ep onto argstk, return new position ap
func push(ep int, argstk *posbuf, ap int) int {
	if ap > EVALSIZE {
		io.Error("macro: argument stack overflow")
	}
	argstk[ap] = ep
	// CHECK: increment ap here? seems no
	return ap + 1
}

// define -- simple string replacement macro processor
func Define() {
	// inlined initdef
	initbuf()
	inithash()

	install(defname, null, DEFTYPE)

	var defn string
	var toktype sttype // type returned by lookup
	for token := gettok(MAXTOK); len(token) > 0; token = gettok(MAXTOK) {
		if !isletter(token[0]) {
			io.Putstr(token, io.STDOUT)
		} else {

			if !lookup(token, &defn, &toktype) {
				io.Putstr(token, io.STDOUT) // undefined
			} else if toktype == DEFTYPE { //denf
				token, defn = getdef(MAXTOK, MAXDEF)
				install(token, defn, MACTYPE)
			} else {
				pbstr(defn) // push replacement onto input
			}
		}
	}
}

// getdef -- get name and definition
func getdef(toksize int, defsize int) (string, string) {
	var c byte
	if getpbc(&c) != LPAREN {
		io.Error("define: missing left paren")
		return "", ""
	}
	token := gettok(toksize)
	if !isletter(token[0]) {
		io.Error("define: non-alphanumeric name")
		return "", ""
	}
	if getpbc(&c) != COMMA {
		io.Error("define: missing comman in define")
		return "", ""
	}
	// got '(name,' so far
	for getpbc(&c) == BLANK {
		// skip blank lines
	}
	putback(c)

	var defn = make([]byte, toksize)
	nlpar := 0
	i := 0
	for nlpar >= 0 {
		if i >= defsize {
			io.Error("define: definition too long")
			return "", ""
		}
		if getpbc(&defn[i]) == io.ENDFILE {
			io.Error("define: missing right paren")
			return "", ""
		}
		if defn[i] == LPAREN {
			nlpar++
		} else if defn[i] == RPAREN {
			nlpar--
		} // else normal character in defn[i]
		i++

	}
	return token, string(defn[0 : i-1])
}

// gettok -- get token for define
// The call c := gettok(token, maxtok)
// copies the next token from the standard input into token.
func gettok(toksize int) string {
	var token = make([]byte, toksize)
	i := 0
	for i < toksize {
		if isalphanum(getpbc(&token[i])) {
			i = i + 1
		} else {
			break
		}
	}
	if i >= toksize {
		io.Error("define: token too long")
	}
	if i > 0 { // some alpha was seen
		putback(token[i]) // put the last symbol, reded after alpha back
		return string(token[0:i])

	}
	// else single non-alphanumeric
	// i == 0, we read space or '\n' or # or ENDFILE
	if token[0] == io.ENDFILE {
		return ""
	}
	return string(token[0:1])
}
