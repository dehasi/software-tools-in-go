package main

// entab -- replaces blanks by tabs and blanks
func entab() {
	var c int8 = 0
	var tabstops [MAXLINE]bool
	settabs(&tabstops)
	col := 0

	for { // repeat
		var newcol = col
		for getc(&c) == BLANK { // collect blanks
			newcol += 1
			if tabpos(newcol, tabstops) {
				putc(TAB)
				col = newcol
			}
		}
		for col < newcol { // output leftover blanks
			putc(BLANK)
			col += 1
		}
		if c != ENDFILE {
			putc(c)
			if c == NEWLINE {
				col = 0
			} else {
				col += 1
			}

		}
		if c == ENDFILE { // until
			return
		}
	}
}
