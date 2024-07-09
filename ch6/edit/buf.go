package edit

import (
	"ch6/io"
	"os"
)

type buftype struct {
	txt  int  // text of line
	mark bool // mark for line
}

var buf [MAXLINES]buftype
var scrout *os.File // scratch input fd
var scrin *os.File  // scratch output fd
var recin int       // next record to read from scrin
var recout int      // next record to write on scrout
var edittemp string // temp file name 'edtemp'

// setbuf (scratch file) -- create scratch file, set up line 0
// initializes the buffer to contain only a valid line zero, and creates a scratch file if necessary
func setbuf() {
	edittemp = "edtemp"
	scrout = io.Mustcreate(edittemp)
	scrin = io.Mustopen(edittemp)
	recout = 1
	recin = 1
	curln = 0
	lastln = 0
}

// clrbuf (scratch file) -- dispose of scratch file
func clrbuf() StCode {
	io.Close(scrin)
	io.Close(scrout)
	io.Remove(edittemp)
	return OK
}

// puttxt (scratch file) -- put text from lin after curln
// copies the text in lin into the buffer immediately after the current line and sets curln to the last line added.
func puttxt(inline string) StCode {
	if lastln < MAXLINES {
		lastln = lastln + 1
		io.Putstr(inline, scrout)
		putmark(lastln, false)
		buf[lastln].txt = recout
		recout += 1
		blkmove(lastln, lastln, curln)
		curln = curln + 1
		return OK
	}
	return ERR
}

// gettxt (scratch file) -- get text from line n into s
// copies the contents of line n into the string s.
func gettxt(n int) string {
	// scopy(buf(n].txt, 1, s, 1)
	if n == 0 {
		return ""
	}
	line := seek(buf[n].txt)
	recin = recin + 1
	return line
}

// blkmove -- move block of lines n1 .. n2 to after n3
// rearranges lines by moving the block of lines n 1 through n2 to after line n3. n3 must not be between n 1 and n2.
func blkmove(n1 int, n2 int, n3 int) {
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
func reverse(n1 int, n2 int) {
	for n1 < n2 {
		temp := buf[n1]
		buf[n1] = buf[n2]
		buf[n2] = temp
		n1 = n1 + 1
		n2 = n2 - 1
	}
}

// places the mark m on line n for global prefix processing.
func putmark(n int, mark bool) {
	buf[n].mark = mark
}

// returns the mark on line n.
func getmark(n int) bool {
	return buf[n].mark
}

// seek (UCB) -- special version of primitive for edit
func seek(recno int) string {
	// cheat: reopen scratch file by name
	fd := io.Mustopen(edittemp)
	for i := 1; i < recno; i++ {
		_, _ = io.Getline(fd, io.MAX_STR)
	}
	line, _ := io.Getline(fd, io.MAX_STR)
	return line
}
