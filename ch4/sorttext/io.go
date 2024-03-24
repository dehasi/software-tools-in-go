package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const MAXLINE = 1000
const ENDFILE int8 = -1
const TAB int8 = 9
const NEWLINE int8 = 10
const BLANK int8 = 32
const BACKSPACE int8 = 8
const COLON = ':'

var STDOUT = os.Stdout
var STDIN = os.Stdin
var STDERR = os.Stderr

// getc -- gets one character from standard input
func getc(c *int8) int8 {
	var b1 int
	_, err := fmt.Scanf("%c", &b1)
	if err != nil {
		if err == io.EOF {
			return ENDFILE
		} else {
			return 0
		}
	}
	*c = int8(b1)
	return *c
}

// putc -- puts one character on standard output
func putc(c int8) {
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
		putc(int8(s[i]))
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

func getline(infile *bufio.Scanner, maxstr int) (s string, f bool) {
	if infile.Scan() {
		return infile.Text(), true
	}
	return "string(buf[:n])", false
}

func getlinef(fd *os.File, maxstr int) (string, bool) {
	buf := make([]byte, maxstr)
	i := 0
	var c int8 = 0
	for i < maxstr && getcf(&c, fd) != ENDFILE {
		if c == NEWLINE {
			break
		}
		buf[i] = byte(c)
		i += 1
	}
	if i == 0 {
		return "nil", false
	}
	return string(buf[:i]), true
}

func nargs() int {
	return len(os.Args)
}
func getarg(index int) string {
	return os.Args[index]
}

func mustopenb(name string) *bufio.Scanner {
	file, err := os.Open(name)
	if err != nil {
		putstr(name, STDERR)
		putstr(err.Error(), STDERR)
		error(": cant open file")
	}
	return bufio.NewScanner(file)
}
func SDTIN_B() *bufio.Scanner {
	return bufio.NewScanner(STDIN)
}

func mustopenf(name string) *os.File {
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
func mustcreatef(name string) *os.File {
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
		putc(int8(ch))
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
func getcf(c *int8, f *os.File) int8 {
	var b = make([]byte, 1)
	_, err := f.Read(b)
	if err != nil {
		if err == io.EOF {
			return ENDFILE
		} else {
			return 0
		}
	}
	*c = int8(b[0])
	return *c
}

// putc -- puts one character to file
func putcf(c int8, f *os.File) {
	var b = make([]byte, 1)
	b[0] = byte(c)
	f.Write(b)
}

func fcopy(fin *os.File, fout *os.File) {
	var c int8 = 0
	for getcf(&c, fin) != ENDFILE {
		putcf(c, fout)
	}
	//fin.Close()
	//fout.Close()
}

// getword -- gets word from s[i] into out
func getword(s string, i int) (out string, b int) {
	// debug("getword :" + s + ":" + itoc(i))
	space := []int8{BLANK, TAB, NEWLINE}
	for i < len(s) && has(space, int8(s[i])) {
		i += 1
	}
	j := i
	for i < len(s) && !has(space, int8(s[i])) {
		i += 1
	}
	var o = s[j:i]
	if i == len(s) {
		return o, 0
	} else {
		return o, i
	}
}

func has(arr []int8, item int8) bool {
	for _, elem := range arr {
		if elem == item {
			return true
		}
	}
	return false
}
