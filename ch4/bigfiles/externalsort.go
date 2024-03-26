package main

import "os"

const MAX_CHARS = 10_000
const MAX_LINES = 300
const MERGE_ORDER = 5

// sort -- external sort of text lines
//
// Gets input from STDIN (we, suppose our bid file is submittend via std in)
// Reads input by MAX_LINES chunks, sorts them and saves sorted results to files with names "1", "2", "3" etc.
// Merges files "1", "2", "3" to one file
// Outputs result into STDOUT
func sort() {
	done := false
	high := 0
	for !done {
		linebuf := make([]string, MAX_LINES)
		linepos := make([]*string, MAX_LINES)
		lines := gtext(linebuf, linepos, STDIN)
		if lines < 0 {
			break
		}
		quick(linepos, lines, linebuf)
		high += 1
		outfile := mustcreate(itoc(high))
		ptext(linepos, lines, linebuf, outfile)
	}
}

func ptext(linepos []*string, lines int, linebuf []string, fd *os.File) {
	for i := 0; i < lines; i++ {
		putstr(*linepos[i], fd)
		putcf(NEWLINE, fd)
	}
}

func quick(linepos []*string, lines int, linebuf []string) {

}

// gtext -- gets text lines into linebuf, and set pointers in linepos
func gtext(linebuf []string, linepos []*string, fd *os.File) int {
	nlines := len(linebuf)
	i := 0
	for {
		line, got := getlinef(fd, MAX_CHARS)
		if !got {
			break
		}
		linebuf[i] = line
		linepos[i] = &linebuf[i]
		i += 1
		if i >= nlines {
			return -1 // to many lines
		}
	}
	return i
}

func main() {
	sort()
}
