package main

import (
	"fmt"
	"io"
)

const MAXLINE = 1000
const ENDFILE int8 = -1
const TAB = 9
const NEWLINE int8 = 10
const BLANK int8 = 32
const BACKSPACE int8 = 8

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
	var s = itoc(nc)
	nd := len(s)
	for i := nd; i < wide; i++ {
		putc(BLANK)
	}
	for i := 0; i < nd; i++ {
		putc(int8(s[i]))
	}
}

// itoc - converts integer n to string
func itoc(n int) string {
	if n < 0 {
		return "-" + itoc(-n)
	}
	if n >= 10 {
		return itoc(n/10) + string(rune('0'+(n%10)))
	}
	return string(rune('0' + (n % 10)))
}
