package main

import "os"

const MAX_CHARS = 10_000
const MAX_LINES = 10 // for testing
const MERGE_ORDER = 5

// sort -- external sort of text lines
//
// Gets input from STDIN (we, suppose our bid file is submittend via std in)
// Reads input by MAX_LINES chunks, sorts them and saves sorted results to files with names "1", "2", "3" etc.
// Merges files "1", "2", "3" to one file
// Outputs result into STDOUT
func sort() {
	high := 0
	linebuf := make([]string, MAX_LINES)
	linepos := make([]*string, MAX_LINES) // maybe use array of integers instead
	for {
		nlines := gtext(linebuf, linepos, STDIN)
		if nlines < 0 {
			break
		}
		quick(linepos, nlines, linebuf)
		high = high + 1
		outfile := makefile(high)
		ptext(linepos, nlines, linebuf, outfile)
		close(outfile)
	}
	low := 1 // it should be 0
	for low < high {
		lim := min(low+MERGE_ORDER-1, high)
		infile := gopen(low, lim)
		high += 1
		outfile := makefile(high)
		merge(infile, lim-low+1, outfile)
		close(outfile)
		gremove(infile, low, lim)
		low = low + MERGE_ORDER
	}
	name := gname(high)
	outfile := mustopenf(name)
	fcopy(outfile, STDOUT)
	close(outfile)
	remove(name)
}

// makefile -- make new file for number n
func makefile(n int) *os.File {
	name := gname(n)
	return mustcreate(name)
}

// gname -- generates unique name for file id n
func gname(n int) string {
	return "stemp" + itoc(n)
}

func gremove(infile []*os.File, f1 int, f2 int) {
	n := f2 - f1 + 1
	for i := 0; i < n; i++ {
		close(infile[i])
		name := gname(f1 + i) // -1
		remove(name)
	}
}

func gopen(f1 int, f2 int) []*os.File {
	n := f2 - f1 + 1
	result := make([]*os.File, n)
	for i := 0; i < n; i++ {
		name := gname(f1 + i) // -1
		result[i] = mustopenf(name)
	}
	return result
}

func merge(infile []*os.File, i int, outfile *os.File) {

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
