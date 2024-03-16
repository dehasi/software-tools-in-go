package main

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func overstrike() {
	const SKIP = BLANK
	const NOSKIP = '+'
	var c int8 = 0
	var col = 0
	for { // repeat
		var newcol = col
		for getc(&c) == BACKSPACE { // eat backspaces
			newcol = max(newcol-1, 1)

			if newcol < col {
				putc(NEWLINE) // start overstrike line
				putc(NOSKIP)
				for i := 0; i < newcol; i++ {
					putc(BLANK)
				}
				col = newcol
			} else if (col == 1) && (c != ENDFILE) {
				putc(SKIP)
			}
			// else: middle of line
			if c != ENDFILE {
				putc(c)
				if c == NEWLINE {
					col = 0
				} else {
					col += 1
				}
			}
		}
		// until
		if c == ENDFILE {
			return
		}
	}
}
