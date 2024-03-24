package main

import (
	"os"
)

const MAXCHARS = 10_000
const MAXLINES = 300

func inmemsort() {

	posbuf := make([]string, MAXLINES)
	res := gtext(posbuf, STDIN)
	if res {
		// shell(posbuf)
		ptext(posbuf, STDOUT)
	}
}

func ptext(strings []string, fd *os.File) {
	n := len(strings)
	for i := 0; i < n; i++ {
		putstr(strings[i], fd)
		putcf(NEWLINE, fd)
	}
}

func gtext(posbuf []string, fd *os.File) bool {
	i := 0
	for {
		line, got := getlinef(fd, MAXCHARS)
		if !got {
			break
		}
		posbuf[i] = line
		i += 1
	}
	return true
}

func main() {

}
