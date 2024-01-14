// complete copy -- to show one possible implementation
package main

import (
	"fmt"
	"io"
	"strconv"
)

const ENDFILE int8 = -1
const NEWLINE int8 = 10

// getc -- get one character from standard input
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

// putc -- put one character on standard output
func putc(c int8) {
	if c == NEWLINE {
		fmt.Println()
	} else {
		fmt.Printf("%c", c)
	}
}

func putdec(nc int, i int) {
	var str = strconv.Itoa(nc)
	for _, ch := range str {
		putc(int8(ch))
	}
}

// charcount -- count characters in standard input
func charcount() {
	var c int8 = 0
	nc := 0
	for getc(&c) != ENDFILE {
		nc += 1
	}
	putdec(nc, 1)
	putc(NEWLINE)
}

// main program
func main() {
	charcount()
}
