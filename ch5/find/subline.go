package main

const DITTO uint8 = '\x01'

// subline -- substitute sub for pat in lin and print
func subline(line, pat, sub string) {
	last_m := -1
	for i := 0; i < len(line); {
		m := amatch(line, i, pat, 0)
		if m > -1 && m != last_m {
			// replace matched text
			putsub(line, i, m, sub)
			last_m = m
		}
		if m == -1 || m == i {
			// no match or null match
			putc(line[i])
			i = i + 1
		} else {
			// skip matched text
			i = m
		}
	}
}

// putsub -- output substitution text
func putsub(line string, from int, to int, sub string) {
	for i := 0; i < len(sub); i++ {
		if sub[i] == DITTO {
			for j := from; j < to; j++ {
				putc(line[j])
			}
		} else {
			putc(sub[i])
		}
	}
}

func getsub(s string) (string, bool) {
	return makesub(s), true
}

// makesub -- make substitution string from arg in sub
// TODO: think how to avoid constant string concatenation
func makesub(arg string) string {
	sub := ""
	for i := 0; i < len(arg); i++ {
		if arg[i] == '&' {
			sub += string(DITTO)
		} else {
			sub += string(esc(arg, i))
		}
	}
	return sub
}
