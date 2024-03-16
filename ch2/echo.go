package main

import (
	"os"
)

func echo() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, arg := range args {
			putc(BLANK)
			for i := range arg {
				putc(int8(arg[i]))
			}
		}
		putc(NEWLINE)
	}
}
