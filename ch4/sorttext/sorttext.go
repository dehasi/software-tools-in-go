package main

import (
	"os"
)

const MAXCHARS = 10_000
const MAXLINES = 300

func inmemsort() {

	linebuf := make([]string, MAXLINES)
	nlines := gtext(linebuf, STDIN)
	if nlines > 0 {
		// shell(posbuf)
		ptext(linebuf, nlines, STDOUT)
	}
}

func ptext(strings []string, nlines int, fd *os.File) {
	for i := 0; i < nlines; i++ {
		putstr(strings[i], fd)
		putcf(NEWLINE, fd)
	}
}

func gtext(linebuf []string, fd *os.File) int {
	i := 0
	for {
		line, got := getlinef(fd, MAXCHARS)
		if !got {
			break
		}
		linebuf[i] = line
		i += 1
	}
	return i
}

func main() {
	inmemsort()

}
