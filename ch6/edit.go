package main

import (
	"ch6/io"
)

type StCode int8

const (
	OK      StCode = 0
	ERR     StCode = 1
	ERRDATA StCode = 2
	ENDDATA StCode = 3 // wtf?
)

var line1 int32 = 0 // first line number
var line2 int32 = 0 // second line number
var nlines int32    // # of line numbers specified
var curln int32     // current line -- value of dot
var lastln int32    // last line -- value of $

// getlist -- get list of line nums at lin[i], increment i
func getlist(lin string, i int32, status StCode) StCode {
	var num int32 = 0
	done := getone(lin, i, num, status) != OK

	for !done {
		line1 = line2
		line2 = num
		nlines = nlines + 1
		if lin[i] == ':' {
			curln = int32(num)
		}
		if lin[i] == ',' || lin[i] == ':' {
			i = i + 1
			done = getone(lin, i, num, status) != OK
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
	return status

}

func getone(lin string, i int32, num int32, status StCode) StCode {
	// istart, mul, pnum int32
	istart := i
	num = 0
	if getnum(lin, i, num, status) == OK {
		for { // repeat

			skipbl(lin, i)
			if lin[i] != '+' && lin[i] != '-' {
				status = ENDDATA
			}
			// until
			if status == OK {
				break
			}
		}
	}
	if num < 0 || num > lastln {
		return ERR
	}
	if i <= istart {
		return ENDDATA
	}
	return OK

}

// getnum -- get single line number component
func getnum(lin string, i int32, num int32, status StCode) StCode {
	status = OK
	skipbl(lin, i)
	panic("unimplemented")
}

// skipbl -- skip blanks and tabs at s[i]
func skipbl(s string, i int32) int32 {
	for s[i] == io.TAB || s[i] == io.BLANK {
		i += 1
	}
	return i
}

// nextln -- get line after n
func nextln(n int32) int32 {
	if n >= lastln {
		return 0
	} else {
		return n + 1
	}
}

// prevln -- get line before n
func prevln(n int32) int32 {
	if n <= 0 {
		return lastln
	} else {
		return n - 1
	}
}

func main() {

}
