package edit

import "ch6/io"

const PLUS uint8 = '+'
const MINUS uint8 = '-'
const PERIOD uint8 = '.'
const COMMA uint8 = ','
const SEMICOL uint8 = ';'

const MAXLINES = 100
const MAXPAT = io.MAX_STR
const DITTO = -1
const CLOSURE uint8 = '*'
const BOL uint8 = '%'
const EOL uint8 = '$'
const ANY uint8 = '?'
const CCL uint8 = '['
const CCLEND uint8 = ']'
const NEGATE uint8 = '^'
const NCCL uint8 = '!'
const LITCHAR = 'c'
const CURLINE uint8 = PERIOD
const LASTLINE uint8 = '$'
const SCAN uint8 = '/'
const BACKSCAN uint8 = '\\'

const ACMD uint8 = 'a'
const CCMD uint8 = 'c'
const DCMD uint8 = 'd'
const ECMD uint8 = 'e'
const EQCMD uint8 = '='
const FCMD uint8 = 'f'
const GCMD uint8 = 'g'
const ICMD uint8 = 'i'
const MCMD uint8 = 'm'
const PCMD uint8 = 'p'
const QCMD uint8 = 'q'
const RCMD uint8 = 'r'
const SCMD uint8 = 's'
const WCMD uint8 = 'w'
const XCMD uint8 = 'x'
