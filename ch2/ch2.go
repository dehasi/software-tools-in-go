package main

const WARNING = '~'

// tabpos -- retursn true is cik is a tab stop
func tabpos(col int, tabstops [1000]bool) bool {
	if col > MAXLINE {
		return true
	} else {
		return tabstops[col]
	}
}

// settabs -- sets initial tab stops
func settabs(tabstops *[MAXLINE]bool) {
	const TAB_SPACE = 4
	for i := 0; i < MAXLINE; i++ {
		tabstops[i] = (i % TAB_SPACE) == 0
	}
}

func main() {
	//entab()
	//overstrike()
	//compress()
	//expand()
	//echo()
	translit()
}
