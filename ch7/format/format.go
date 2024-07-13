package format

import "ch7/io"

const CMD uint8 = '.'
const HUGE int = 42

type CmdType int8

const (
	BP      CmdType = 1
	BR      CmdType = 2
	CE      CmdType = 3
	FI      CmdType = 4
	FO      CmdType = 5
	HE      CmdType = 6
	IND     CmdType = 7
	LS      CmdType = 8
	NF      CmdType = 9
	PL      CmdType = 10
	RM      CmdType = 11
	SP      CmdType = 12
	TI      CmdType = 13
	UL      CmdType = 14
	UNKNOWN CmdType = 0
)

var fill bool // fill if true; init=true
var lsval int // current line spacing; init=1
var spval int // # of lines to space }
var inval int // current indent; >= 0; init=O
var rmval int // right margin; init=PAGEWIDTH=60
var tival int // current temporary indent; init=O
var ceval int // # of lines to center; init=O
var ulval int // # of lines to underline; init=O

// maybe different type
var header string
var footer string

func Format() {
	initfmt()
	for inbuf, result := io.Getline(io.STDIN, io.MAX_STR); result; inbuf, result = io.Getline(io.STDIN, io.MAX_STR) {
		if inbuf[0] == CMD {
			command(inbuf)
		} else {
			text(inbuf)
		}
	}
}

// command -- perform formatting command
func command(buf string) {
	//	cmd : cmdtype;
	// argtype, spval, val integer;
	val := 0
	spval := 0
	argtype := 0
	cmd := getcmd(buf)
	if cmd != UNKNOWN {
		val = getval(buf, argtype)
	}
	switch cmd { // TODO: figure our why sometimes we have 'break' right after 'begin'
	case FI:
		fill = true
	case NF:
		fill = false
	case BR:
		break
	case LS:
		setparam(lsval, val, argtype, 1, 1, HUGE)
	case CE:
		setparam(ceval, val, argtype, 1, 0, HUGE)

	case UL:
		setparam(ulval, val, argtype, 1, 0, HUGE)

	case HE:
		gettl(buf, header)
	case FO:
		gettl(buf, header)

	}
}

// break -- end current filled line
func breakk() {
	panic("unimplemented")
}

// page -- get to top of new page
func page() {
	panic("unimplemented")
}

// space -- space n lines or to bottom of page
func space(n int) {
	panic("unimplemented")
}

// gettl -- copy title from buf to ttl
func gettl(buf, ttl string) {
	panic("unimplemented")
}

// puttl -- put out title line with optional page number
func puttl(but string, pageno int) {
	panic("unimplemented")
}

// put -- put out line with proper spacing and indenting
func put(buf string) {
	panic("unimplemented")
}

// puthead -- put out page header
func puthead() {
	panic("unimplemented")
}

// putfoot -- put out page footer
func putfoot() {
	panic("unimplemented")
}

func setparam(param, val, argtype, defval, minval, maxval int) {
	panic("unimplemented")
}

func getcmd(buf string) CmdType {
	panic("unimplemented")
}

// getval -- evaluate optional numeric argument
func getval(buf string, argtype int) int {
	panic("unimplemented")
}

// text -- process text lines (interim version 1)
func text(inbuf string) {
	put(inbuf)
}

func initfmt() {

}
