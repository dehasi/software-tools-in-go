package main

import (
	"fmt"
	"io"
	"os"
)

const MAXLINE = 1000
const MAX_STR = 1000    // max symbols per line
const ENDFILE uint8 = 0 // -1?
const TAB uint8 = 9
const NEWLINE uint8 = 10
const BLANK uint8 = 32
const BACKSPACE uint8 = 8
const COLON = ':'

var STDOUT = os.Stdout
var STDIN = os.Stdin
var STDERR = os.Stderr

// getc -- gets one character from standard input
func getc(c *uint8) uint8 {
	var b1 int
	_, err := fmt.Scanf("%c", &b1)
	if err != nil {
		if err == io.EOF {
			return ENDFILE
		} else {
			return 0
		}
	}
	*c = uint8(b1)
	return *c
}

// putc -- puts one character on standard output
func putc(c uint8) {
	if c == NEWLINE {
		fmt.Println()
	} else {
		fmt.Printf("%c", c)
	}
}

// putdec -- puts number digits to  standard output
func putdec(nc int, wide int) {
	var s = itoc(nc)
	nd := len(s)
	for i := nd; i < wide; i++ {
		putc(BLANK)
	}
	for i := 0; i < nd; i++ {
		putc(uint8(s[i]))
	}
}

// itoc - converts integer n to string
func itoc(n int) string {
	if n < 0 {
		return "-" + itoc(-n)
	}
	if n >= 10 {
		return itoc(n/10) + string(rune('0'+(n%10)))
	}
	return string(rune('0' + (n % 10)))
}

// ctoi - converts string to integer
func ctoi(c string) int {
	n := 0
	for i := 0; i < len(c); i++ {
		n *= 10
		n += int(c[i] - '0')
	}
	return n
}

func putstr(s string, f *os.File) {
	_, err := f.WriteString(s)
	if err != nil {
		return
	}
}

func getline(fd *os.File, maxstr int) (string, bool) {
	buf := make([]byte, maxstr)
	i := 0
	var c uint8 = 0
	for i < maxstr && getcf(&c, fd) != ENDFILE {
		buf[i] = c
		i += 1
		if c == NEWLINE {
			break
		}
	}
	if i == 0 {
		return "nil", false
	}
	if c == ENDFILE {
		i += 1
		buf[i] = NEWLINE
	}
	return string(buf[:i]), true
}

func nargs() int {
	return len(os.Args)
}
func getarg(index int) string {
	return os.Args[index]
}

func mustopen(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		putstr(name, STDERR)
		putstr(err.Error(), STDERR)
		error(": cant open file")
	}
	return file
}

func create(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		return nil
	}
	return file
}

func mustcreate(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		putstr(name, STDERR)
		putstr(err.Error(), STDERR)
		error(": cant open file")
	}
	return file
}
func atef(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		putstr(name, STDERR)
		putstr(err.Error(), STDERR)
		error(": cant create file: " + name)
	}
	return file
}

func close(fd *os.File) {
	if fd == STDIN || fd == STDOUT || fd == STDERR {
		return
	}
	fd.Close()
}

func remove(filename string) {
	os.Remove(filename)
}

func error(message string) {
	for _, ch := range message {
		putc(uint8(ch))
	}
	putc(NEWLINE)
	os.Exit(42)
}

// message -- prints message string to the output
func message(message string) {
	putstr(message, STDOUT)
	putc(NEWLINE)
}

// debug -- prints debug message string to the output
func debug(m string) {
	//message(m)
}

// getcf -- gets one character from file
func getcf(c *uint8, f *os.File) uint8 {
	var b = make([]byte, 1)
	_, err := f.Read(b)
	if err != nil {
		if err == io.EOF {
			return ENDFILE
		} else {
			error("ERROR: " + err.Error())
			return ENDFILE // this line is unreachable
		}
	}
	*c = b[0]
	return *c
}

// putcf -- puts one character to file
func putcf(c uint8, f *os.File) {
	var b = make([]byte, 1)
	b[0] = c
	f.Write(b)
}

func fcopy(fin *os.File, fout *os.File) {
	var c uint8 = 0
	for getcf(&c, fin) != ENDFILE {
		putcf(c, fout)
	}
	//fin.Close()
	//fout.Close()
	close(fin)
	close(fout)
}

// getword -- gets word from s[i] into out
func getword(s string, i int) (out string, b int) {
	// debug("getword :" + s + ":" + itoc(i))
	space := []uint8{BLANK, TAB, NEWLINE}
	for i < len(s) && has(space, uint8(s[i])) {
		i += 1
	}
	j := i
	for i < len(s) && !has(space, uint8(s[i])) {
		i += 1
	}
	var o = s[j:i]
	if i == len(s) {
		return o, 0
	} else {
		return o, i
	}
}

func has(arr []uint8, item uint8) bool {
	for _, elem := range arr {
		if elem == item {
			return true
		}
	}
	return false
}
