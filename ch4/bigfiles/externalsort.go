package main

import "os"

const MAX_CHARS = 10_000 // MAX_STR* MAX_LINES => max size of the whole text
const MAX_STR = 100      // max symbols per line
const MAX_LINES = 10     // for testing
const MERGE_ORDER = 5

// sort -- external sort of text lines
//
// Gets input from STDIN (we, suppose our bid file is submittend via std in)
// Reads input by MAX_LINES chunks, sorts them and saves sorted results to files with names "1", "2", "3" etc.
// Merges files "1", "2", "3" to one file
// Outputs result into STDOUT
func sort() {
	high := 0
	linebuf := make([]uint8, MAX_CHARS) // holds all chars of the text
	linepos := make([]int, MAX_LINES)   // holds indexes, where a line begins in linebuf
	for {
		nlines := gtext(linebuf, linepos, STDIN)
		if nlines <= 0 {
			break
		}
		quick(linepos, nlines, linebuf)
		high = high + 1
		outfile := makefile(high)
		ptext(linepos, nlines, linebuf, outfile)
		close(outfile)
	}

	low := 1 // it should be 0?
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

// gtext -- gets text lines into linebuf, and set pointers in linepos
func gtext(linebuf []uint8, linepos []int, fd *os.File) int {
	nlines := 0
	nextpos := 0

	for {
		temp, done := getlinef(fd, MAX_STR)
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
		if nlines >= MAX_LINES {
			return -1 // to many lines
		}
	}
	return nlines
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

// charpos = 1...MAX_CHARS => int
// type
// charpos = 1 .. MAXCHARS;
// charbuf = array [1 .. MAXCHARS] of character;
// posbuf = array [1 .. MAXLINES] of charpos;
// pos = O.. MAXLINES;
// fdbuf = array [1 .. MERGEORDER] of filedesc;
// merge -- merges infile[1] ... infile[nf] onto outfile
func merge(infile []*os.File, nf int, outfile *os.File) {
	// technically we cam ,ake them global variables and reuse.
	// after colleciting input and putting it into files they are not used
	linebuf := make([]uint8, MAX_CHARS) // holds all chars of the text
	linepos := make([]int, MAX_LINES)   // holds indexes, where a line begins in linebuf
	j := 0
	for i := 0; i < nf; i++ { //get one line from each file
		temp, read := getlinef(infile[i], MAX_CHARS)
		if read { // if we read the line => file stll not empty
			lbp := (i)*MAX_STR + 1 // room for longest, (i-1)
			sccopy(temp, linebuf, lbp)
			linepos[i] = lbp
			j = j + 1
		}
	}
	nf = j                      // some infile[] will be eventually empty
	quick(linepos, nf, linebuf) // make initial heap
	for nf > 0 {
		lbp := linepos[0] // lowest line
		temp := cscopy(linebuf, lbp)
		putstr(temp, outfile)
		i := lbp/MAX_STR + 1 // compute file index
		temp, got := getlinef(infile[i], MAX_STR)
		if got {
			sccopy(temp, linebuf, lbp)
		} else { // one less input file
			linepos[0] = linepos[nf-1]
			nf = nf - 1
		}
		reheap(linepos, nf, linebuf)
	}

}

func reheap(linepos []int, nf int, linebuf []uint8) {

}

// cscopy -- copy ch[i] - array of chars - to string
func cscopy(cb []uint8, i int) string {
	start := i
	for cb[i] != ENDSTR {
		i++
	}
	end := i

	// Strings in Go are immutable, let's make a string from a slice of the array
	return string(cb[start : end+1])
}

// sccopy -- copy string s to cb[i] - array of chars
func sccopy(temp string, cb []uint8, i int) {
	for j := 0; temp[j] != ENDSTR; {
		cb[i] = temp[j]
		i++
		j++
	}
	cb[i] = ENDSTR
}

func main() {
	sort()
}
