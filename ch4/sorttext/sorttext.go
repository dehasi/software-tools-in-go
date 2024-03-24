package main

import (
	"os"
)

const MAXCHARS = 10_000
const MAXLINES = 300

func inmemsort() {

	posbuf := make([]string, MAXLINES)
	nlines := gtext(posbuf, STDIN)
	if nlines > 0 {
		// shell(posbuf)
		ptext(posbuf, nlines, STDOUT)
	}
}

func ptext(strings []string, nlines int, fd *os.File) {
	println("ptext, n=" + itoc(nlines))
	for i := 0; i < nlines; i++ {
		putstr(strings[i], fd)
		putcf(NEWLINE, fd)
	}
}

func gtext(posbuf []string, fd *os.File) int {
	i := 0
	for {
		line, got := getlinef(fd, MAXCHARS)
		if !got {
			break
		}
		posbuf[i] = line
		i += 1
	}
	return i
}

func main() {
	inmemsort()

}
