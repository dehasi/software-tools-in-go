package edit

import (
	"strings"
)

// After impelemting p 6.10 Scratch Files, this buffer won't be used
// Keep for history

type inMemBuftype struct {
	txt  string // text of line
	mark bool   // mark for line
}

var buf_inmem [MAXLINES]inMemBuftype

// setbuf (in memory) -- initialize line storage buffer
// initializes the buffer to contain only a valid line zero, and creates a scratch file if necessary
func setbuf_inmem() {
	buf_inmem[0].txt = ""
	curln = 0
	lastln = 0
}

// clrbuf (in memory) -- initialize for new file
// discards the scratch file, if one is used.
func clrbuf_inmem() StCode {
	// in memory, nothing to do
	return OK
}

// puttxt (in memory) -- put text from lin after curln
// copies the text in lin into the buffer immediately after the current line and sets curln to the last line added.
func puttxt_inmem(inline string) StCode {
	// println("puttxt", "inline:", inline)
	if lastln < MAXLINES {
		lastln = lastln + 1
		buf_inmem[lastln].txt = strings.Clone(inline)
		putmark(lastln, false)
		blkmove(lastln, lastln, curln)
		curln = curln + 1
		return OK
	}
	return ERR
}

// gettxt (in memory) -- get text from line n into s
// copies the contents of line n into the string s.
func gettxt_inmem(n int) string {
	// println("gettxt", "n:", n)
	// scopy(buf(n].txt, 1, s, 1)
	return strings.Clone(buf_inmem[n].txt)
}

// blkmove -- move block of lines n1 .. n2 to after n3
// rearranges lines by moving the block of lines n 1 through n2 to after line n3. n3 must not be between n 1 and n2.
func blkmove_inmem(n1 int, n2 int, n3 int) {
	if n3 < n1-1 {
		reverse(n3+1, n1-1)
		reverse(n1, n2)
		reverse(n3+1, n2)
	} else if n3 > n2 {
		reverse(n1, n2)
		reverse(n2+1, n3)
		reverse(n1, n3)
	}
}

// reverse -- reverse buf[n1] ... buf[n2]
func reverse_inmem(n1 int, n2 int) {
	for n1 < n2 {
		temp := buf[n1]
		buf[n1] = buf[n2]
		buf[n2] = temp
		n1 = n1 + 1
		n2 = n2 - 1
	}
}

// places the mark m on line n for global prefix processing.
func putmark_inmem(n int, mark bool) {
	buf[n].mark = mark
	// panic("unimplemented")
}

// returns the mark on line n.
func getmark_inmem(n int) bool {
	return buf[n].mark
	// panic("unimplemented")
}
