package main

import "os"

const MAX_CHARS = 10_000
const MAX_LINES = 300
const MERGE_ORDER = 5

// sort -- external sort of text lines
func sort() {
	done := false
	for !done {

	}
}

// gtext -- gets text lines into linebuf, and set pointers in linepos
func gtext(linebuf []string, nlines int, linepos []*string, fd *os.File) int {
	i := 0
	for i < nlines {
		line, got := getlinef(fd, MAX_CHARS)
		if !got {
			break
		}
		linebuf[i] = line
		linepos[i] = &linebuf[i]
		i += 1
		if i >= MAX_LINES {
			return -1 // to many lines
		}
	}
	return i
}

func main() {

}
