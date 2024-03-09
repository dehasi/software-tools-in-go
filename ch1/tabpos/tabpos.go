package tabpos

const MAXLINE = 1000

// Tabpos -- retursn true is cik is a tab stop
func Tabpos(col int, tabstops [1000]bool) bool {
	if col > MAXLINE {
		return true
	} else {
		return tabstops[col]
	}
}
