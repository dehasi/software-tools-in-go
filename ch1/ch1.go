// complete copy -- to show one possible implementation
package main

import (
	"ch1/tabpos"
	"fmt"
	"io"
	"strconv"
)

const MAXLINE = 1000
const ENDFILE int8 = -1
const TAB = 9
const NEWLINE int8 = 10
const BLANK int8 = 32

// getc -- gets one character from standard input
func getc(c *int8) int8 {
	var b1 int
	_, err := fmt.Scanf("%c", &b1)
	if err != nil {
		if err == io.EOF {
			return ENDFILE
		} else {
			return 0
		}
	}
	*c = int8(b1)
	return *c
}

// putc -- puts one character on standard output
func putc(c int8) {
	if c == NEWLINE {
		fmt.Println()
	} else {
		fmt.Printf("%c", c)
	}
}

// putdec -- puts number digits to  standard output
func putdec(nc int, wide int) {
	var str = strconv.Itoa(nc)
	for _, ch := range str {
		putc(int8(ch))
	}
}

// copy -- copies input to output
func copy() {
	var c int8 = 0
	for getc(&c) != ENDFILE {
		putc(c)
	}
}

// charcount -- counts characters in standard input
func charcount() {
	var c int8 = 0
	nc := 0
	for getc(&c) != ENDFILE {
		nc += 1
	}
	putdec(nc, 1)
	putc(NEWLINE)
}

// linecount -- counts lines in standard input
func linecount() {
	var c int8 = 0
	nl := 0
	for getc(&c) != ENDFILE {
		if c == NEWLINE {
			nl += 1
		}
	}
	putdec(nl, 1)
	putc(NEWLINE)
}

// wordcount -- counts words in standard input
func wordcount() {
	var c int8 = 0
	nw := 0
	in_word := false
	for getc(&c) != ENDFILE {
		if c == NEWLINE || c == TAB || c == BLANK {
			in_word = false
		} else if !in_word {
			in_word = true
			nw += 1
		}
	}
	putdec(nw, 1)
	putc(NEWLINE)
}

// detab -- converts tabs to equivalent number of blanks
func detab() {
	var c int8 = 0
	var tabstops [MAXLINE]bool
	settabs(tabstops)
	col := 0
	for getc(&c) != ENDFILE {
		if c == TAB {
			// imitation do-while loop
			putc(BLANK)
			col += 1
			for tabpos.Tabpos(col, tabstops) {
				putc(BLANK)
				col += 1
			}
		} else if c == NEWLINE {
			putc(c)
			col = 0
		} else {
			putc(c)
			col += 1
		}
	}
}

// main program
func main() {
	//copy()
	//charcount()
	//linecount()
	//wordcount()
	detab()
}
