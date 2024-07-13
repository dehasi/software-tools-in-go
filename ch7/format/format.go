package format

import "ch7/io"

// fmtcons -- constants for format
const CMD uint8 = '.'
const PAGENUM uint8 = '#'
const PAGEWIDTH int = 60
const PAGELEN int = 66
const HUGE int = 10000

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

// page parameters
var curpage int   // current output page number; init=O
var newpage int   // next output page number; init=1
var lineno int    // next line to be printed; init=O
var plval int     // page length in lines; init=PAGELEN=66
var m1val int     // margin before and including header
var m2val int     // margin after header
var m3val int     // margin after last text line
var m4val int     // bottom margin, including footer
var bottom int    // last line on page, =plval-m3val-m4val
var header string // top of page title; init=NEWLINE
var footer string // bottom of page title; init=NEWLINE

// global parameters
var fill bool // fill if true; init=true
var lsval int // current line spacing; init=1
var spval int // # of lines to space }
var inval int // current indent; >= 0; init=O
var rmval int // right margin; init=PAGEWIDTH=60
var tival int // current temporary indent; init=O
var ceval int // # of lines to center; init=O
var ulval int // # of lines to underline; init=O

// output area
var outp int      // last char pos in outbuf; init=O
var outw int      // width of text in outbuf; init=O
var outwds int    // number of words in outbuf; init=O
var outbuf string // lines to be filled collect here
var dir int       // 0..1 direction for blank padding
var inbuf string  // input line

func Format() {
	initfmt()
	result := false
	for inbuf, result = io.Getline(io.STDIN, io.MAX_STR); result; inbuf, result = io.Getline(io.STDIN, io.MAX_STR) {
		if inbuf[0] == CMD {
			command(inbuf)
		} else {
			text(inbuf)
		}
	}
	page() // flush last output, if any
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
		val = getval(buf, &argtype)
	}
	switch cmd {
	case FI:
		breakk()
		fill = true

	case NF:
		breakk()
		fill = false

	case BR:
		breakk()

	case LS:
		setparam(&lsval, val, argtype, 1, 1, HUGE)

	case CE:
		breakk()
		setparam(&ceval, val, argtype, 1, 0, HUGE)

	case UL:
		setparam(&ulval, val, argtype, 1, 0, HUGE)

	case HE:
		gettl(buf, header)

	case FO:
		gettl(buf, header)

	case BP:
		page()
		setparam(&curpage, val, argtype, curpage+1, -HUGE, HUGE)
		newpage = curpage

	case SP:
		setparam(&spval, val, argtype, 1, 0, HUGE)
		space(spval)

	case IND:
		setparam(&inval, val, argtype, 0, 0, rmval-1)

	case RM:
		setparam(&rmval, val, argtype, PAGEWIDTH, inval+tival+1, HUGE)

	case TI:
		breakk()
		setparam(&tival, val, argtype, 0, -HUGE, rmval)

	case PL:
		setparam(&plval, val, argtype, PAGELEN, m1val+m2val+m3val+m4val+1, HUGE)
		bottom = plval - m3val - m4val

	case UNKNOWN: // ignore
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

// setparam -- set parameter and check range
func setparam(param *int, val, argtype, defval, minval, maxval int) {
	if argtype == int(io.NEWLINE) { // defaulted
		*param = defval
	} else if argtype == '+' { // relative +
		*param += val
	} else if argtype == '-' { // relative -
		*param -= val
	} else { // absolute
		*param = val
	}
	*param = min(*param, maxval)
	*param = max(*param, minval)
}

// getcmd -- decode command type
func getcmd(buf string) CmdType {
	cmd := buf[1:2]
	switch cmd {
	case "fi":
		return FI
	case "nf":
		return NF
	case "br":
		return BR
	case "ls":
		return LS
	case "bp":
		return BP
	case "sp":
		return SP
	case "in":
		return IND
	case "rm":
		return RM
	case "ti":
		return TI
	case "ce":
		return CE
	case "ul":
		return UL
	case "he":
		return HE
	case "fo":
		return FO
	case "pl":
		return PL

	default:
		return UNKNOWN
	}
}

// getval -- evaluate optional numeric argument
// TODO: check if I need to return argtype
func getval(buf string, argtype *int) int {
	i := 0
	// skip over command nam
	for !io.IsBlank(buf[i]) {
		i++
	}
	i = io.Skipbl(buf, i)
	*argtype = int(buf[i])
	if *argtype == '+' || *argtype == '-' {
		i = i + 1
	}
	return io.Ctoi(buf[i:])
}

// text -- process text lines (interim version 1)
func text(inbuf string) {
	put(inbuf)
}

// initfmt -- set format parameters to default values
func initfmt() {
	fill = true
	dir = 0
	inval = 0
	rmval = PAGEWIDTH
	tival = 0
	lsval = 1
	spval = 0
	ceval = 0
	ulval = 0
	lineno = 0
	curpage = 0
	newpage = 1
	plval = PAGELEN
	m1val = 3
	m2val = 2
	m3val = 2
	m4val = 3
	bottom = plval - m3val - m4val
	header = string(io.NEWLINE) // initial titles
	footer = string(io.NEWLINE)

	outp = 0
	outw = 0
	outwds = 0
}
