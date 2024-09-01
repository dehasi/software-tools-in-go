package macro

import "ch8/io"

const BUFSIZE int = 500 // size of pushback buffer
var buf [BUFSIZE]byte   // for putback, starts with 1
var bp = 0              // next available cahracter; init=0

func initbuf() {
	bp = 0
}

// getpbc -- get a (possibly pushed back) character
func getpbc(c *byte) byte {
	if bp > 0 {
		*c = buf[bp]
	} else {
		bp = 1
		buf[bp] = io.Getc(c)
	}
	if *c != io.ENDFILE { // delete from buf
		bp--
	}
	return *c
}

// putback -- push character back onto input
func putback(c byte) {
	if bp > BUFSIZE {
		io.Error("too many characters pushed back")
	}
	bp++
	buf[bp] = c
}

// pbstr -- push string back onto input
func pbstr(s string) {
	for i := len(s) - 1; i >= 0; i-- {
		putback(s[i])
	}
}

// pbnum -- convert number to string, push back on input
func pbnum(n int) {
	pbstr(io.Itoc(n))
}
