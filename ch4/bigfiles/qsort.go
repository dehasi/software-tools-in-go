package main

const ENDSTR uint8 = 10 // Go dosn't have '\0' concept as C, let's use '\n'

// quick -- quicksort for lines
func quick(linepos []int, nlines int, linebuf []uint8) {
	rquick(0, nlines-1, linepos, linebuf)
}

// rquick -- recursive quicksort
func rquick(lo int, hi int, linepos []int, linebuf []uint8) {
	if lo < hi {
		i := lo
		j := hi
		pivline := linepos[j] // pivot line
		for /*repeat*/ {
			for i < j && cmp(linepos[i], pivline, linebuf) <= 0 {
				i += 1
			}
			for j > i && cmp(linepos[j], pivline, linebuf) >= 0 {
				j -= 1
			}
			if i < j { // out of order pair
				exchange(linepos, i, j)
			}
			if /*until*/ i >= j {
				break
			}
		}
		// move pivot to i
		exchange(linepos, i, hi)

		// I don't understand this optimization
		if (i - lo) < (hi - i) {
			rquick(lo, i-1, linepos, linebuf)
			rquick(i+1, hi, linepos, linebuf)
		} else {
			rquick(i+1, hi, linepos, linebuf)
			rquick(lo, i-1, linepos, linebuf)
		}
	}
}

// cmp -- compare linebuf[lp1] with linebuf[lp2]
func cmp(i int, j int, linebuf []uint8) int {
	for linebuf[i] == linebuf[j] &&
		linebuf[i] != ENDSTR {
		i = i + 1
		j = j + 1
	}
	// get the first not equal symbol
	if linebuf[i] == linebuf[j] {
		return 0
	} else if linebuf[i] == ENDSTR { // 1st is shorter
		return -1
	} else if linebuf[j] == ENDSTR { // 2st is shorter
		return 1
	} else if linebuf[i] < linebuf[j] {
		return -1
	} else {
		return 1
	}
}

// exchange -- exchanges linepos[lp1] with linepos[lp2]; divergents from the original
func exchange(linepos []int, lp1 int, lp2 int) {
	tmp := linepos[lp1]
	linepos[lp1] = linepos[lp2]
	linepos[lp2] = tmp
}
