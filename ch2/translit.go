package main

import (
	"os"
	"strconv"
)

const NEGATE = '^'
const ENDSTR = '\n'
const ESCAPE = '@'
const DASH = '-'
const MAXSTR = 50
const NOT_FOUND = -1

func translit() {
	// check len(os.Args) and end with error if necessary
	if len(os.Args) < 2 {
		error("usage: translit from to")
	}

	allbut := os.Args[1][0] == NEGATE
	var fromset = dodash(os.Args[1])

	if allbut {
		fromset = fromset[1:]
	}

	if len(fromset) > MAXSTR {
		error("translit: 'from' set too large")
	}

	var toset = " os.Args[2]"
	if len(os.Args) != 3 {
		toset = string(ENDSTR)
	} else {
		toset = dodash(os.Args[2])
		if len(toset) > MAXSTR {
			error("translit: 'to' set too large")
		}
	}

	if len(fromset) < len(toset) {
		error("translit: 'from' shorter than 'to'")
	}

	lastto := len(toset)
	squash := (len(fromset) > lastto) || allbut
	var c int8 = 0

	for { // REPEAT
		i := xindex(fromset, getc(&c), allbut, lastto)
		if squash && (i >= lastto) && (lastto >= 0) { //mb lastto >= 0
			putc(int8(toset[lastto-1]))
			// REPEAT
			i = xindex(fromset, getc(&c), allbut, lastto)
			for i >= lastto { // UNTIL
				i = xindex(fromset, getc(&c), allbut, lastto)
			}
		}
		if c != ENDFILE {
			if i >= 0 && lastto >= 0 { // translate
				putc(int8(toset[i]))
			} else if i == NOT_FOUND { // copy
				putc(c)
			} else {
				//delete
			}
		}
		if c == ENDFILE { // UNTIL
			return
		}
	}
}

// index -- returns first index of 'c'
func index(fromset string, c int8) int {
	for i, ch := range fromset {
		if c == int8(ch) {
			return i
		}
	}
	return NOT_FOUND
}

// xindex -- conditionally inverts value from index
func xindex(inset string, c int8, allbut bool, lastto int) int {
	if c == ENDFILE {
		return NOT_FOUND
	} else if !allbut {
		return index(inset, c)
	} else if index(inset, c) >= 0 {
		return NOT_FOUND
	} else { // allbut = true, c is not in inset
		return lastto + 1
	}
}

// dodash - expands set at src[i] into dest[j], stop at delim
func dodash(src string) string {
	var result = ""

	for i := 0; i < len(src); i++ {

		if src[i] == ESCAPE {
			result += strconv.Itoa(int(esc(src, i)))
			i++ // we need to increment, if it's last it's ok. we end form the loop
		} else if src[i] != DASH {
			result += string(src[i])
		} else if i == len(src)-1 { // last element is just dash
			result += string(DASH)
		} else if src[i-1] < src[i+1] {
			for k := src[i-1] + 1; k < src[i+1]; k++ {
				result += string(k)
			}
		} else {
			result += strconv.Itoa(int(DASH))
		}
	}

	return result
}

// esc -- maps s[i] into escaped character, increment i
func esc(s string, i int) int8 {
	if s[i] != ESCAPE {
		return int8(s[i])
	} else if i+1 == len(s) { // @ not special at end
		return ESCAPE
	} else {
		i++
		if s[i] == 'n' {
			return NEWLINE
		} else if s[i] == 't' {
			return TAB
		} else {
			return int8(s[i])
		}
	}
}
func error(message string) {
	for _, ch := range message {
		putc(int8(ch))
	}
	putc(NEWLINE)
	os.Exit(42)
}
