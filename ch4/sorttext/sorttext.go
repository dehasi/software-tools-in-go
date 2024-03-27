package main

import (
	"os"
)

const MAXCHARS = 10_000
const MAXSTR = 1000
const MAXLINES = 300
const ENDSTR uint8 = 10 // Go dosn't have '\0' concept as C, let's use '\n'

// inmemsort -- sorts text lines in memory
func inmemsort() {

	linebuf := make([]uint8, MAXCHARS) // holds all chars of the text
	linepos := make([]int, MAXLINES)   // holds indexes, where a line begins in linebuf

	nlines := gtext(linebuf, linepos, STDIN)
	if nlines > 0 {
		//shell(linepos, nlines, linebuf)
		quick(linepos, nlines, linebuf)
		ptext(linepos, nlines, linebuf, STDOUT)
	}
}

// gtext -- gets text lines into linebuf, and set pointers in linepos
func gtext(linebuf []uint8, linepos []int, fd *os.File) int {
	nlines := 0
	nextpos := 0
	for {
		temp, done := getlinef(fd, MAXSTR)
		if !done {
			break
		}
		linepos[nlines] = nextpos
		slen := len(temp)
		for i := 0; i < slen; i++ {
			linebuf[nextpos+i] = temp[i]
		}
		linebuf[nextpos+slen] = ENDSTR
		nlines = nlines + 1
		nextpos = nextpos + slen + 1
		if nlines >= MAXLINES {
			return -1 // to many lines
		}
	}
	return nlines
}

// shell -- ascending Shell sort for lines
func shell(linepos []int, nlines int, linebuf []uint8) {
	for gap := nlines / 2; gap > 0; gap /= 2 {
		for i := gap; i < nlines; i++ {
			for j := i - gap; j >= 0; j = j - gap {
				jq := j + gap
				if cmp(linepos[j], linepos[jq], linebuf) <= 0 {
					break
				}
				exchange(linepos, j, jq)
			}
		}
	}
}

// cmp -- compare linebuf[lp1] with linebuf[lp2]
func cmp(i int, j int, linebuf []uint8) int {
	for linebuf[i] == linebuf[j] &&
		linebuf[i] != ENDSTR {
		i = i + 1
		j = j + 1
	}
	// get the first not equal symbol
	if linebuf[i] == linebuf[j] {
		return 0
	} else if linebuf[i] == ENDSTR { // 1st is shorter
		return -1
	} else if linebuf[j] == ENDSTR { // 2st is shorter
		return 1
	} else if linebuf[i] < linebuf[j] {
		return -1
	} else {
		return 1
	}
}

// exchange -- exchanges linepos[lp1] with linepos[lp2]; divergents from the original
func exchange(linepos []int, lp1 int, lp2 int) {
	tmp := linepos[lp1]
	linepos[lp1] = linepos[lp2]
	linepos[lp2] = tmp
}

// exchange -- exchanges linebuf[lp1] with linebuf[lp2]
//func exchange(lp1 int, lp2 int) { }

// ptext -- outputs text lines from linebuf
func ptext(linepos []int, nlines int, linebuf []uint8, outfile *os.File) {
	for i := 0; i < nlines; i++ {
		j := linepos[i]
		for linebuf[j] != ENDSTR {
			putcf(int8(linebuf[j]), outfile)
			j += 1
		}
		putcf(NEWLINE, outfile)
	}
}

// quick -- quicksort for lines
func quick(linepos []int, nlines int, linebuf []uint8) {
	rquick(0, nlines-1, linepos, linebuf)
}

// rquick -- recursive quicksort
func rquick(lo int, hi int, linepos []int, linebuf []uint8) {
	if lo < hi {
		i := lo
		j := hi
		pivline := linepos[j] // pivot line
		for /*repeat*/ {
			for i < j && cmp(linepos[i], pivline, linebuf) <= 0 {
				i += 1
			}
			for j > i && cmp(linepos[j], pivline, linebuf) >= 0 {
				j -= 1
			}
			if i < j { // out of order pair
				exchange(linepos, i, j)
			}
			if /*until*/ i >= j {
				break
			}
		}
		// move pivot to i
		exchange(linepos, i, hi)

		// I don't understand this optimization
		if (i - lo) < (hi - i) {
			rquick(lo, i-1, linepos, linebuf)
			rquick(i+1, hi, linepos, linebuf)
		} else {
			rquick(i+1, hi, linepos, linebuf)
			rquick(lo, i-1, linepos, linebuf)
		}
	}
}
func main() {
	inmemsort()
}
