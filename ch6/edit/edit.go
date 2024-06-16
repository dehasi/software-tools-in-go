package edit

import (
	"ch6/find"
	"ch6/io"
)

type StCode int8

const (
	OK      StCode = 0
	ERR     StCode = 1
	ENDDATA StCode = 2
)

var line1 int = 0 // first line number
var line2 int = 0 // second line number
var nlines int    // # of line numbers specified
var curln int     // current line -- value of dot
var lastln int    // last line -- value of $
var pat string    // pattern

func Edit() {

	for line, ok := io.Getline(io.STDIN, io.MAXLINE); ok; line, ok = io.Getline(io.STDIN, io.MAXLINE) {
		// println("process line: ", line)
		i := 0
		cursave := curln
		i, status := getlist(line, i)
		if status == OK {
			if ckglob(line, i, status) == OK {
				status = doglob(line, i, cursave, status)
			} else if status != ERR {
				status = docmd(line, i, false)
			} // else error, do nothing
		} else if status == ERR {
			io.Putstr("?", io.STDOUT)
		} else if status == ENDDATA {
			break
		}
	}
}

// setDefault -- set defaulted line numbers, original name 'default'
func setDefault(def1 int, def2 int) StCode {

	if nlines == 0 {
		line1 = def1
		line2 = def2
		return OK
	}

	if (line1 > line2) || (line1 < 0) {
		return ERR
	}
	return OK
}

// doprint -- print lines n1 through n2
func doprint(n1 int, n2 int) StCode {
	// println("doprint", "n1: ", n1, "n2:", n2)
	if n1 < 0 {
		return ERR
	}
	for i := n1; i <= n2; i++ {
		line := gettxt(i)
		io.Putstr(line, io.STDOUT)
	}
	curln = n2
	return OK
}

// docmd -- handle all commands except globals
// The false argument to docmd says that it is not being called from within a global prefix
func docmd(lin string, i int, glob bool) StCode {
	// println("docmd", "lin", lin, "i", i, "lin[i]", lin[i])
	status := ERR
	pflag := false // may be set by d, m, s
	switch lin[i] {
	case PCMD:
		// println("PCMD")
		if lin[i+1] == io.NEWLINE {
			if setDefault(curln, curln) == OK {
				status = doprint(line1, line2)
			}
		}
	case io.NEWLINE:
		if nlines == 0 {
			line2 = nextln(curln)
		}
		status = doprint(line2, line2)
	case QCMD:
		if (lin[i+1] == io.NEWLINE) && nlines == 0 && (!glob) {
			status = ENDDATA
		}
	case ACMD:
		// println("ACMD")
		if lin[i+1] == io.NEWLINE {
			status = append(line2, glob)
		}
	default:
		status = ERR
	}
	if status == OK && pflag {
		status = doprint(curln, curln)
	}
	// println("docmd return status", status)
	return status
}

// append -- append lines after "line"
func append(line int, glob bool) StCode {
	// println("append", "line:", line, "glob:", glob)

	if glob {
		return ERR
	}
	curln = line
	stat := OK
	done := false

	for !done && stat == OK {
		inline, ok := io.Getline(io.STDIN, io.MAX_STR)
		if !ok {
			stat = ENDDATA
		} else if inline[0] == PERIOD && inline[1] == io.NEWLINE {
			done = true
		} else if puttxt(inline) == ERR {
			stat = ERR
		}
	}
	return stat
}

func doglob(line string, i, cursave int, status StCode) StCode {
	panic("unimplemented")
}

// ckglob looks for g/.../ or x/.../;
// if either is found, ckglob marks the lines for processing by doglob,
func ckglob(line string, i int, status StCode) StCode {
	return ENDDATA
}

// getlist -- get list of line nums at lin[i], increment i
func getlist(lin string, i int) (int, StCode) {
	// println("getlist", "lin", lin, "i", i)
	line2 = 0
	nlines = 0
	num, i, status := getone(lin, i)
	done := status != OK
	for !done { // if not ERR, jump in
		line1 = line2
		line2 = num
		nlines = nlines + 1
		if lin[i] == SEMICOL {
			curln = num
		}
		if lin[i] == COMMA || lin[i] == SEMICOL {
			i = i + 1
			num, i, status = getone(lin, i)
			done = status != OK
		} else {
			done = true // TODO break
		}
	}
	nlines = min(nlines, 2)
	if nlines == 0 {
		line2 = curln
	}
	if nlines <= 1 {
		line1 = line2
	}
	if status != ERR {
		status = OK
	}

	// println("getlist return", "i:", i, "status:", status, "num:", num)
	return i, status

}

// getone -- get one line number expression
func getone(lin string, i int) (int, int, StCode) {
	// println("getone", "lin:", lin, "i:", i)
	istart := i
	var mul int = 0
	num, i, status := getnum(lin, i)
	if status == OK {
		for { // repeat + or - terms

			i = skipbl(lin, i)
			if lin[i] != PLUS && lin[i] != MINUS {
				status = ENDDATA
			} else {
				if lin[i] == PLUS {
					mul += 1
				} else {
					mul -= 1
				}
				i = i + 1
				var pnum = 0
				pnum, i, status = getnum(lin, i)
				if status == OK {
					num = num + mul*pnum
				}
				if status == ENDDATA {
					status = ERR
				}
			}
			// until status <> OK
			if status != OK {
				break
			}
		}
	}
	if num < 0 || num > lastln {
		return -1, -1, ERR
	}
	if i <= istart {
		return 1, -1, ENDDATA
	}
	return num, i, OK

}

// getnum -- get single line number component
// evaluates one number in a line number expression, where a number is either an integer, . (dot), $, or a context search
func getnum(lin string, i int) (int, int, StCode) {
	// we expet that 'num' will hold a result
	// we expect that i will also be changed (incresed)
	// so ideally we have to return i, nm, status
	status := OK
	num := 0
	i = skipbl(lin, i)
	if io.Isdigit(lin[i]) {
		i, num = io.Ctoi2(lin, i)
		i = i - 1 // move back; to be advanced at end
	} else if lin[i] == CURLINE {
		num = curln
	} else if lin[i] == LASTLINE {
		num = lastln
	} else if lin[i] == SCAN || lin[i] == BACKSCAN {
		if optpat(lin, i) == ERR { // build pattern
			status = ERR
		} else {
			status = patscan(lin[i], num)
		}
	} else {
		status = ENDDATA
	}

	if status == OK {
		i = i + 1 // next character to be examined
		return num, i, OK
	}
	return -1, -1, status
}

// patscan -- find next occurrence of pattern after line n
func patscan(way byte, n int) StCode {

	// n = curln // do we need it? it always be after curln
	patscanSt := ERR
	done := false
	line := ""
	for {
		if way == SCAN {
			n = nextln(n)
		} else {
			n = prevln(n)
		}
		line = gettxt(n)
		if find.Match(line, pat) {
			patscanSt = OK
			done = true
		}
		// until n == curln || done
		if n == curln || done {
			break
		}
	}
	return patscanSt
}

// optpat -- get optional pattern from lin[i], increment i
func optpat(lin string, i int) StCode {
	// not sure what we need to return besides i, maybe pat?
	// or pat should be global like lastln?
	n := len(lin)
	if n == i || n == i+1 {
		i = 0
	} else if lin[i+1] == lin[i] { // repeated delimiter
		//  leave existing pattern alone
		i = i + 1
	} else {
		pat = find.Makepat(lin, i+1, lin[i])
		if pat == "" {
			i = 0
		}
	}
	if i == 0 {
		pat = ""
		return ERR
	}
	return OK
}

// skipbl -- skip blanks and tabs at s[i]
func skipbl(s string, i int) int {
	for s[i] == io.TAB || s[i] == io.BLANK {
		i += 1
	}
	return i
}

// nextln -- get line after n
func nextln(n int) int {
	if n >= lastln {
		return -1 // 0 in original, arrays in pascal starts form 1
	} else {
		return n + 1
	}
}

// prevln -- get line before n
func prevln(n int) int {
	if n < 0 {
		return lastln
	} else {
		return n - 1
	}
}
