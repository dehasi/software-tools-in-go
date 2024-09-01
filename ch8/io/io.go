package io

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

// Getc -- gets one character from standard input
func Getc(c *uint8) uint8 {
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

// Putc -- puts one character on standard output
func Putc(c uint8) {
	if c == NEWLINE {
		fmt.Println()
	} else {
		fmt.Printf("%c", c)
	}
}

// Putdec -- puts number digits to  standard output
func Putdec(nc int, wide int) {
	var s = Itoc(nc)
	nd := len(s)
	for i := nd; i < wide; i++ {
		Putc(BLANK)
	}
	for i := 0; i < nd; i++ {
		Putc(uint8(s[i]))
	}
}

// itoc - converts integer n to string
func Itoc(n int) string {
	if n < 0 {
		return "-" + Itoc(-n)
	}
	if n >= 10 {
		return Itoc(n/10) + string(rune('0'+(n%10)))
	}
	return string(rune('0' + (n % 10)))
}

// Isdigit -- checks is a char is a digit
func Isdigit(b byte) bool {
	return ('0' <= b) && (b <= '9')
}

// Ctoi2 - converts string to integer from i, and as possible
func Ctoi2(c string, i int) (int, int) {
	n := 0
	for ; i < len(c) && Isdigit(c[i]); i++ {
		n *= 10
		n += int(c[i] - '0')
	}
	return i, n
}

// Ctoi - converts string to integer
func Ctoi(c string) int {
	n := 0
	for i := 0; i < len(c); i++ {
		n *= 10
		n += int(c[i] - '0')
	}
	return n
}

func Putstr(s string, f *os.File) {
	_, err := f.WriteString(s)
	if err != nil {
		return
	}
}

func Getline(fd *os.File, maxstr int) (string, bool) {
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

func NArgs() int {
	return len(os.Args)
}
func GetArg(index int) string {
	return os.Args[index]
}

func Mustopen(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		Putstr(name, STDERR)
		Putstr(err.Error(), STDERR)
		Error(": cant open file")
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

func Mustcreate(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		Putstr(name, STDERR)
		Putstr(err.Error(), STDERR)
		Error(": cant open file")
	}
	return file
}
func atef(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		Putstr(name, STDERR)
		Putstr(err.Error(), STDERR)
		Error(": cant create file: " + name)
	}
	return file
}

func Close(fd *os.File) {
	if fd == STDIN || fd == STDOUT || fd == STDERR {
		return
	}
	fd.Close()
}

func Remove(filename string) {
	os.Remove(filename)
}

func Error(message string) {
	for _, ch := range message {
		Putc(uint8(ch))
	}
	Putc(NEWLINE)
	os.Exit(42)
}

// message -- prints message string to the output
func message(message string) {
	Putstr(message, STDOUT)
	Putc(NEWLINE)
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
			Error("ERROR: " + err.Error())
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
	Close(fin)
	Close(fout)
}

// getword -- gets word from s[i] into out
func Getword(s string, i int) (out string, ni int) {
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

// skipbl -- skip blanks and tabs at s[i]
func Skipbl(s string, i int) int {
	// As Go strings don't have EOL marker, so I need to be creative
	for i < len(s) && s[i] != NEWLINE && (s[i] == TAB || s[i] == BLANK) {
		i += 1
	}
	return i
}

// IsBlank == checks is c is tab or space or newline
func IsBlank(c uint8) bool {
	return c == BLANK || c == TAB || c == NEWLINE
}

func has(arr []uint8, item uint8) bool {
	for _, elem := range arr {
		if elem == item {
			return true
		}
	}
	return false
}
