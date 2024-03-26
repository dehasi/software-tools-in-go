package main

// quick -- quicksort for lines
func quick(linepos []*string, nlines int, linebuf []string) {
	rquick(linepos, 0, nlines-1)
}

// rquick -- recursive quicksort
func rquick(linepos []*string, lo int, hi int) {
	if lo >= hi {
		return
	}
	i := lo
	j := hi
	piv := *linepos[hi]
	for i < j {
		for i < j && *linepos[i] <= piv {
			i++
		}
		for i < j && *linepos[j] >= piv {
			j--
		}
		if i < j {
			tmp := linepos[j]
			linepos[j] = linepos[i]
			linepos[i] = tmp
		}
	}
	// move pivot to i
	tmp := linepos[hi]
	linepos[hi] = linepos[i]
	linepos[i] = tmp

	rquick(linepos, lo, i-1)
	rquick(linepos, i+1, hi)
}
