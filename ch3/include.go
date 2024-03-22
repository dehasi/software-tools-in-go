package main

import (
	"bufio"
)

// include -- replaces #include "file" by contents of file
func include() {
	finclude(bufio.NewScanner(STDIN))
}

// finclude -- includes file desc f
func finclude(f *bufio.Scanner) {
	for {
		line, e := getline(f, MAXLINE)
		if !e {
			return
		}
		str, loc := getword(line, 0)
		if str != "#include" {
			putstr(str+"\n", STDOUT)
		} else {
			str, loc = getword(line, loc)
			str = str[1 : len(str)-1] // remove quotes
			f1 := mustopenb(str)
			finclude(f1)
			// close f
		}
	}
}
