package main

// initializes the buffer to contain only a valid line zero, and creates a scratch file if necessary
func setbuf() StCode {
	panic("unimplemented")
}

// discards the scratch file, if one is used.
func clrbuf() StCode {
	panic("unimplemented")
}

// copies the text in lin into the buffer immediately after the current line and sets curln to the last line added.
func puttxt(inline string) StCode {
	panic("unimplemented")
}

// copies the contents of line n into the string s.
func gettxt(n int) string {
	// I think it should open file and get text at line 'n'
	panic("unimplemented")
}

// rearranges lines by moving the block of lines n 1 through n2 to after line n3. n3 must not be between n 1 and n2.
func blkmove(n1 int, n2 int, n3 int) {
	panic("unimplemented")
}

// places the mark m on line n for global prefix processing.
func putmark(n int, m int) {
	panic("unimplemented")
}

// returns the mark on line n.
func getmark(n int) {
	panic("unimplemented")
}
