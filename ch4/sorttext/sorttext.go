package main

import (
	"os"
)

const MAXCHARS = 10_000
const MAXLINES = 300

// inmemsort -- sorts text lines in memory }
func inmemsort() {

	linebuf := make([]string, MAXLINES)
	linepos := make([]*string, MAXLINES)

	nlines := gtext(linebuf, linepos, STDIN)
	if nlines > 0 {
		//shell(linepos, nlines)
		quick(linepos, nlines)
		putstr("SORTED:\n", STDOUT)
		ptext(linepos, nlines, STDOUT)
		putstr("ORIGINAL:\n", STDOUT)
		ptexto(linebuf, nlines, STDOUT)
	}
}

// gtext -- gets text lines into linebuf, and set pointers in linepos
func gtext(linebuf []string, linepos []*string, fd *os.File) int {
	i := 0
	for {
		line, got := getlinef(fd, MAXCHARS)
		if !got {
			break
		}
		linebuf[i] = line
		linepos[i] = &linebuf[i]
		i += 1
		if i >= MAXLINES {
			return -1 // to many lines
		}
	}
	return i
}

// shell -- ascending Shell sort for lines
func shell(linebuf []*string, nlines int) {
	for gap := nlines / 2; gap > 0; gap /= 2 {
		for i := gap; i < nlines; i++ {
			for j := i - gap; j >= 0; j = j - gap {
				jq := j + gap
				if *linebuf[j] <= *linebuf[jq] {
					break
				}
				tmp := linebuf[j]
				linebuf[j] = linebuf[jq]
				linebuf[jq] = tmp
			}
		}
	}
}

// ptext -- outputs text lines from linepos
func ptext(linepos []*string, nlines int, fd *os.File) {
	for i := 0; i < nlines; i++ {
		putstr(*linepos[i], fd)
		putcf(NEWLINE, fd)
	}
}

// ptexts -- outputs text lines from linebuf
func ptexto(linebuf []string, nlines int, fd *os.File) {
	for i := 0; i < nlines; i++ {
		putstr(linebuf[i], fd)
		putcf(NEWLINE, fd)
	}
}

// quick -- quicksort for lines
func quick(linepos []*string, nlines int) {
	rquick(0, nlines-1, linepos)
}

// rquick -- recursive quicksort
func rquick(lo int, hi int, linepos []*string) {
	if lo < hi {
		i := lo
		j := hi
		pivline := linepos[j] // pivot line
		for /*repeat*/ {
			for i < j && *linepos[i] <= *pivline {
				i += 1
			}
			for j > i && *linepos[j] >= *pivline {
				j -= 1
			}
			if i < j { // out of order pair
				tmp := linepos[i]
				linepos[i] = linepos[j]
				linepos[j] = tmp
			}
			if /*until*/ i >= j {
				break
			}
		}
		// move pivot to i
		tmp := linepos[i]
		linepos[i] = linepos[hi]
		linepos[hi] = tmp

		// I don't understand this optimization
		if (i - lo) < (hi - i) {
			rquick(lo, i-1, linepos)
			rquick(i+1, hi, linepos)
		} else {
			rquick(i+1, hi, linepos)
			rquick(lo, i-1, linepos)
		}
	}
}
func main() {
	inmemsort()
}
