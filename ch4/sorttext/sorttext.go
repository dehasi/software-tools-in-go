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
		shell(linebuf, nlines)
		ptext(linebuf, nlines, STDOUT)
	}
}

func shell(linebuf []string, nlines int) {
	for gap := nlines / 2; gap > 0; gap /= 2 {
		for i := gap; i < nlines; i++ {
			for j := i - gap; j >= 0; j = j - gap {
				jq := j + gap
				if linebuf[j] <= linebuf[jq] {
					break
				}
				tmp := linebuf[j]
				linebuf[j] = linebuf[jq]
				linebuf[jq] = tmp
			}
		}
	}
}

func ptext(linebuf []string, nlines int, fd *os.File) {
	for i := 0; i < nlines; i++ {
		putstr(linebuf[i], fd)
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
