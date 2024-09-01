package macro

// global type degifiniotns
type charpos = int // 1..MAXCHARS
type charbuf [MAXCHARS]byte

type sttype int // symbol table types

const (
	DEFTYPE sttype = iota // for the built-in "define" a
	MACTYPE               // for a macro name
	EXPRTYPE
	SUBTYPE
	LENTYPE
	IFTYPE
	CHQTYPE
)

// eval stk
const MAXPOS = 500
const CALLSIZE = MAXPOS
const ARGSIZE = MAXPOS
const EVALSIZE = MAXCHARS
const ARGFLAG = '$'

type posbuf = [MAXPOS]charpos
type pos = int // O.. MAXPOS;

var callstk posbuf           // call stack
var cp pos                   // current call stack position
var typestk [CALLSIZE]sttype // type
var plev [CALLSIZE]int       // paren level
var argstk posbuf            // argument stack for this call
var ap pos                   // current argument position
var evalstk charbuf          // evaluation stack
var ep charpos               // first character unused in evalstk
