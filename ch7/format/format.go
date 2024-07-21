package format

import (
	"ch7/io"
)

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
	IN      CmdType = 7
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
	// println("command", "val:", val)
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
		header = gettl(buf)

	case FO:
		footer = gettl(buf)

	case BP:
		page()
		setparam(&curpage, val, argtype, curpage+1, -HUGE, HUGE)
		newpage = curpage

	case SP:
		setparam(&spval, val, argtype, 1, 0, HUGE)
		space(spval)

	case IN:
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
	//panic("unimplemented")
	if outp > 0 {
		put(outbuf + "\n")
	}
	outp = 0
	outw = 0
	outwds = 0
}

// page -- get to top of new page
func page() {
	breakk()
	if lineno > 0 && lineno <= bottom {
		skip(bottom + 1 - lineno)
		putfoot()
	}
	lineno = 0
}

// space -- space n lines or to bottom of page
func space(n int) {
	if lineno <= bottom {
		if lineno <= 0 {
			puthead()
		}
		skip(min(n, bottom+1-lineno))
		lineno = lineno + n
		if lineno > bottom {
			putfoot()
		}
	}
}

// gettl -- copy title from buf to ttl
// The title is assumed to begin with the first non-blank character,
// but a leading apostrophe or quote is stripped off, to permit a title to begin with blanks,
// for example, to right-justify it.
func gettl(buf string) string {
	i := 0
	// skip command name
	for !io.IsBlank(buf[i]) {
		i++
	}
	i = io.Skipbl(buf, i) // find argumen
	if buf[i] == '\'' || buf[i] == '"' {
		i++ // strip leading quote
	}
	return buf[i:]
}

// puttl -- put out title line with optional page number
func puttl(buf string, pageno int) {
	for i := 0; i < len(buf); i++ {
		if buf[i] == PAGENUM {
			io.Putdec(pageno, 1)
		} else {
			io.Putc(buf[i])
		}
	}
}

// put -- put out line with proper spacing and indenting
func put(buf string) {
	if lineno <= 0 || lineno > bottom {
		puthead()
	}
	for i := 1; i < inval+tival; i++ { // indenting
		io.Putc(io.BLANK)
	}
	tival = 0
	io.Putstr(buf, io.STDOUT)
	skip(min(lsval-1, bottom-lineno))
	lineno += lsval
	if lineno > bottom {
		putfoot()
	}
}

// leadbl -- delete leading blanks, set tival
func leadbl(buf string) string {
	breakk()
	i := 0
	for ; buf[i] == io.BLANK; i++ {
	}
	if buf[i] != io.NEWLINE {
		tival = tival + i - 1
	}
	return buf[i:]
}

// skip -- produces n empty lines (NEWLINE only) if n is positive, and does nothing if n is less than one
func skip(n int) {
	for i := 1; i <= n; i++ {
		io.Putc(io.NEWLINE)
	}
}

// puthead -- put out page header
func puthead() {
	curpage = newpage
	newpage = newpage + 1
	if m1val > 0 {
		skip(m1val - 1)
		puttl(header, curpage)
	}
	skip(m2val)
	lineno = m1val + m2val + 1
}

// putfoot -- put out page footer
func putfoot() {
	skip(m3val)
	if m4val > 0 {
		puttl(footer, curpage)
		skip(m4val - 1)
	}
}

// setparam -- set parameter and check range
func setparam(param *int, val, argtype, defval, minval, maxval int) {
	// println("setparam:", val, argtype, defval, minval, maxval)
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
	// println("setparam:", "param:", *param)
}

// getcmd -- decode command type
func getcmd(buf string) CmdType {
	cmd := buf[1:3]
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
		return IN
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
func getval(buf string, argtype *int) int {
	i := 0
	// skip over command name
	for !io.IsBlank(buf[i]) {
		i++
	}
	i = io.Skipbl(buf, i)
	*argtype = int(buf[i])
	if *argtype == '+' || *argtype == '-' {
		i = i + 1
	}
	_, n := io.Ctoi2(buf, i)
	return n
}

// center -- center a line by setting tival
func center(buf string) {
	tival = max((rmval+tival-width(buf))/2, 0)
}

const Reset = "\033[0m"
const Underline = "\033[4m"

// underln -- underline a line
func underln(buf string, size int) string {
	return Underline + buf[0:len(buf)-1] + Reset + "\n"
}

// text -- process text lines (interim version 2)
func text(inbuf string) {
	if inbuf[0] == io.BLANK || inbuf[0] == io.NEWLINE {
		inbuf = leadbl(inbuf)
	}
	if ulval > 0 {
		inbuf = underln(inbuf, io.MAX_STR)
		ulval--
	}
	if ceval > 0 {
		center(inbuf)
		put(inbuf)
		ceval = ceval - 1
	} else if inbuf[0] == io.NEWLINE { // all blank line
		put(inbuf)
	} else if !fill { // unfilled text
		put(inbuf)
	} else { // filled text
		for wordbuf, i := io.Getword(inbuf, 0); i != 0; wordbuf, i = io.Getword(inbuf, i) {
			putword(wordbuf)
		}
	}
}

// putword -- put word in outbuf; does margin justification
func putword(wordbuf string) {
	// , nextra, int
	w := width(wordbuf)
	last := len(wordbuf) + outp + 1 // new end of outbuf
	llval := rmval - tival - inval
	if outp > 0 && ((outw+w > llval) || (last >= io.MAX_STR)) {
		last = last - outp // remember end of wordbuf
		nextra := llval - outw + 1
		if (nextra > 0) && (outwds > 1) {
			outbuf = spread(outbuf, nextra)
			outp = outp + nextra
		}
		breakk() // flush previous line
	}

	if outp > 0 {
		outbuf += " " + wordbuf
	} else {
		outbuf = wordbuf
	}
	outp = last
	outw = outw + w + 1 // 1 for blank
	outwds = outwds + 1

}

// spread -- spread words to justify right margin
// intput 'a b c', 2 => 'a  b  c'
func spread(buf string, nextra int) string {
	println("spread", "buf:", buf, ",nextra:", nextra)
	if (nextra > 0) && (outwds > 1) {
		dir = 1 - dir // reverse previous direction
		nholes := outwds - 1
		i := outp - 1
		j := min(io.MAX_STR-2, i+nextra) //room for NEWLINE and ENDSTR ???
		println("j:", j, "len(buf):", len(buf))
		newStr := make([]uint8, j+1)
		for i < j {
			newStr[j] = buf[i]
			if buf[i] == io.BLANK {
				var nb = 0
				if dir == 0 {
					nb = (nextra-1)/nholes + 1
				} else {
					nb = nextra / nholes
				}
				nextra = nextra - nb
				nholes = nholes - 1
				for nb > 0 {
					j = j - 1
					newStr[j] = io.BLANK
					nb = nb - 1
				}
			}
			i--
			j--
		}
		newStr[0] = buf[0]
		return string(newStr)
	} else {
		return buf
	}

}

// width -- compute width of character string
func width(buf string) int {
	w := 0
	for i := 0; i < len(buf); i++ {
		if buf[i] == '\b' { // BACKSPACE
			w--
		} else if buf[i] != io.NEWLINE {

			w++
		}
	}
	return w
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
