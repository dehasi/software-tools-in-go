package macro

import "ch8/io"

// eval -- expand args i .. j: do built-in or push back def
func eval(argstk *posbuf, td sttype, i, j int) {
	// argno, k, t : integer;
	// temp string;
	t := argstk[i]
	if td == DEFTYPE {
		dodef(argstk, i, j)
	} else if td == EXPRTYPE {
		doexpr(argstk, i, j)
	} else if td == SUBTYPE {
		dosub(argstk, i, j)
	} else if td == IFTYPE {
		doif(argstk, i, j)
	} else if td == LENTYPE {
		dolen(argstk, i, j)
	} else if td == CHQTYPE {
		dochq(argstk, i, j)
	} else { // process normal macro as before
		k := t
		for evalstk[k] != ENDSTR {
			k++
		}
		k-- // last character of defn
		for k > t {
			if evalstk[k-1] != ARGFLAG {
				putback(evalstk[k])
			} else {
				argno := int(evalstk[k] - '0')
				if argno >= 0 && argno < j-i {
					temp := cscopy(evalstk[:], argstk[i+argno+1])
					pbstr(temp)
				}
				k-- // skip over $
			}
			k--
		}
		if k == t {
			putback(evalstk[k]) // do last character
		}
	}
}

// dodef -- install definition in table
func dodef(argstk *posbuf, i, j int) {
	if j-i > 2 {
		temp1 := cscopy(evalstk[:], argstk[i+2])
		temp2 := cscopy(evalstk[:], argstk[i+3])
		install(temp1, temp2, MACTYPE)
	}
}

// doexpr -- evaluate arithmetic expressions
func doexpr(argstk *posbuf, i, j int) {
	if j-i > 2 {
		temp := cscopy(evalstk[:], argstk[i+2])
		junk := 0 // 0?
		pbnum(expr(temp, &junk))
	}
}

// expr -- recursive expression evaluation
func expr(s string, i *int) int {
	v := term(s, i)

	for t := gnbchar(s, i); t == '+' || t == '-'; t = gnbchar(s, i) {
		*i++
		if t == '+' {
			v = v + term(s, i)
		} else {
			v -= term(s, i)
		}
	}
	return v
}

// term -- evaluate term of arithmetic expression
func term(s string, i *int) int {
	v := factor(s, i)

	for t := gnbchar(s, i); t == '*' || t == '/' || t == '%'; t = gnbchar(s, i) {
		*i++
		if t == '*' {
			v *= factor(s, i)
		} else if t == '/' {
			v /= factor(s, i)
		} else {
			v %= factor(s, i)
		}
	}
	return v
}

// factor -- evaluate factor of arithmetic expression
func factor(s string, i *int) int {
	if gnbchar(s, i) == LPAREN {
		*i++
		f := expr(s, i)
		if gnbchar(s, i) == RPAREN {
			*i++
		} else {
			io.Error("macro: missing paren in expr")
		}
		return f
	} else {
		ii, n := io.Ctoi2(s, *i)
		*i = ii
		return n
	}
}

// gnbchar -- get next non-blank character
func gnbchar(s string, i *int) byte {
	for *i < len(s) && io.IsBlank(s[*i]) {
		*i++
	}
	if *i >= len(s) {
		return '\x00'
	}
	return s[*i]
}

// dosub -- select substring
func dosub(argstk *posbuf, i, j int) {
	nc := MAXTOK
	if j-i >= 3 {
		if j-i < 4 {
			nc = MAXTOK
		} else {
			temp := cscopy(evalstk[:], argstk[i+4])
			k := 0
			nc = expr(temp, &k)
		}
		temp := cscopy(evalstk[:], argstk[i+3])
		junk := 0
		start := expr(temp, &junk) - 1         //first char
		str := cscopy(evalstk[:], argstk[i+2]) // str to sub
		var subresult string
		if nc >= len(str) {
			subresult = str[start:]
		} else {
			subresult = str[start:(start + nc)]
		}
		pbstr(subresult)
	}
}

// doif -- select one of two arguments
func doif(argstk *posbuf, i, j int) {
	if j-i >= 4 {
		temp1 := cscopy(evalstk[:], argstk[i+2])
		temp2 := cscopy(evalstk[:], argstk[i+3])
		var temp3 string = ""
		if temp1 == temp2 {
			temp3 = cscopy(evalstk[:], argstk[i+4])
		} else if j-i >= 5 {
			temp3 = cscopy(evalstk[:], argstk[i+5])
		} else {
			temp3 = ""
		}
		pbstr(temp3)
	}
}

// dolen -- return length of argument
func dolen(argstk *posbuf, i, j int) {
	if j-i > 1 {
		temp := cscopy(evalstk[:], argstk[i+2])
		pbnum(len(temp))
	} else {
		pbnum(0)
	}
}

// dochq -- change quote characters
func dochq(argstk *posbuf, i, j int) {
	temp := cscopy(evalstk[:], argstk[i+2])
	n := len(temp)
	if n <= 0 { // reset
		lquote = '`'
		rquote = '\''
	} else if n == 1 {
		lquote = rune(temp[0])
		rquote = '\''
	} else {
		lquote = rune(temp[0])
		rquote = rune(temp[1])
	}
}
