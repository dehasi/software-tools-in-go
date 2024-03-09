package main

// settabs -- sets initial tab stops
func settabs(tabstops [MAXLINE]bool) {
	const TAB_SPACE = 4
	for i := 0; i < MAXLINE; i++ {
		tabstops[i] = (i % TAB_SPACE) == 0
	}
}
