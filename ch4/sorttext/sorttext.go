package main

import (
	"os"
)

const MAXCHARS = 10_000
const MAXLINES = 300
const ENDSTR uint8 = 10 // Go dosn't have '\0' concept as C, let's use '\n'

// inmemsort -- sorts text lines in memory
func inmemsort() {

	linebuf := make([]uint8, MAXLINES) // holds all chars of the text
	linepos := make([]int, MAXLINES)   // holds indexes, where a line begins in linebuf

	nlines := gtext(linebuf, linepos, STDIN)
	if nlines > 0 {
		shell(linepos, nlines, linebuf)
		//quick(linepos, nlines)
		ptext(linepos, nlines, linebuf, STDOUT)
	}
}

// gtext -- gets text lines into linebuf, and set pointers in linepos
func gtext(linebuf []uint8, linepos []int, fd *os.File) int {
	nlines := 0
	nextpos := 0
	for {
		temp, done := getlinef(fd, MAXCHARS)
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

// exchange -- exchanges linebuf[lp1] with linebuf[lp2]
func exchange(lp1 int, lp2 int) {
	// looks strange
}

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
//
//	func quick(linepos []*string, nlines int) {
//		rquick(0, nlines-1, linepos)
//	}
//
// // rquick -- recursive quicksort
//
//	func rquick(lo int, hi int, linepos []*string) {
//		if lo < hi {
//			i := lo
//			j := hi
//			pivline := linepos[j] // pivot line
//			for /*repeat*/ {
//				for i < j && *linepos[i] <= *pivline {
//					i += 1
//				}
//				for j > i && *linepos[j] >= *pivline {
//					j -= 1
//				}
//				if i < j { // out of order pair
//					tmp := linepos[i]
//					linepos[i] = linepos[j]
//					linepos[j] = tmp
//				}
//				if /*until*/ i >= j {
//					break
//				}
//			}
//			// move pivot to i
//			tmp := linepos[i]
//			linepos[i] = linepos[hi]
//			linepos[hi] = tmp
//
//			// I don't understand this optimization
//			if (i - lo) < (hi - i) {
//				rquick(lo, i-1, linepos)
//				rquick(i+1, hi, linepos)
//			} else {
//				rquick(i+1, hi, linepos)
//				rquick(lo, i-1, linepos)
//			}
//		}
//	}
func main() {
	inmemsort()
}
