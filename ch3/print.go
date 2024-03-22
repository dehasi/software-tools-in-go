package main

import (
	"bufio"
)

// print --prints files with haedings
func print() {

	if nargs() == 1 {
		fprint("", SDTIN_B())
	} else {
		for i := 1; i < nargs(); i++ {
			name := getarg(i)
			fin := mustopenb(name)
			fprint(name, fin)
		}
	}
}

// fprint - prints file "name" from fin
func fprint(name string, fin *bufio.Scanner) {
	const MARGIN1 = 2
	const MARGIN2 = 2
	const BOTTOM = 64
	const PAGELEN = 66

	skip(MARGIN1)
	pageno := 1
	head(name, pageno)
	skip(MARGIN2)
	lineno := MARGIN1 + MARGIN2 + 1
	for {
		line, readed := getline(fin, MAXLINE)
		if !readed {
			break
		}
		if lineno == 0 {
			skip(MARGIN1)
			pageno += 1
			head(name, pageno)
			skip(MARGIN2)
			lineno = MARGIN1 + MARGIN2 + 1
		}
		putstr(line+"\n", STDOUT)
		lineno += 1
		if lineno >= BOTTOM {
			skip(PAGELEN - lineno)
			lineno = 0
		}
	}
	if lineno >= BOTTOM {
		skip(PAGELEN - lineno)
	}
}

// head -- prints top of page header
func head(name string, pageno int) {
	const PAGE = " Page "
	putstr(name, STDOUT)
	putstr(PAGE, STDOUT)
	putdec(pageno, 1)
	putc(NEWLINE)
}

// skip -- outpus n blank lines
func skip(n int) {
	for ; n > 0; n-- {
		putc(NEWLINE)
	}
}
