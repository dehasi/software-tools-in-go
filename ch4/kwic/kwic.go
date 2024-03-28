package main

const MAXOUT = 80
const MIDDLE = 40
const FOLD uint8 = '$'

// kwic -- makes keyword in context index
func kwic() {
	for {
		buf, read := getline(STDIN, MAX_STR)
		if !read {
			return
		}
		putrot(buf)
	}
}

// putrot -- creates lines with keyword at front
func putrot(buf string) {
	n := len(buf)
	i := 0
	for i < n && buf[i] != NEWLINE { // Go doesn't have END_STR, and doesn't need because we know the len of the string
		if isalphanum(buf[i]) {
			rotate(buf, i)
			for isalphanum(buf[i]) {
				i += 1
			}
		}
		i += 1
	}
}

// rotate -- outputs rotated line
func rotate(buf string, n int) {
	lenb := len(buf)
	for i := n; i < lenb && buf[i] != NEWLINE; i++ {
		putc(buf[i])
	}
	putc(FOLD)
	for i := 0; i < n; i++ {
		putc(buf[i])
	}
	putc(NEWLINE)
}

// isalphanum -- checks if ch is [0-9A-Za-z]
func isalphanum(ch uint8) bool {
	return ('0' <= ch && ch <= '9') ||
		('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z')
}
func main() {
	if nargs() == 1 { // if no args run kwic
		kwic()
	} else { // else run unrorate, we don't care about cli args
		unrotate()
	}
}

// unrotate -- unrotate lines rotated by kwic
func unrotate() {
	const dist = 2
	for inbuf, read := getline(STDIN, MAX_STR); read; inbuf, read = getline(STDIN, MAX_STR) {
		outbuf := make([]uint8, MAXOUT)
		for i := 0; i < MAXOUT; i++ {
			outbuf[i] = BLANK
		}
		f := index(inbuf, FOLD)
		// copy the part after FOLD for left
		j := MIDDLE - dist
		for i := len(inbuf) - 1; i > f; i-- {
			if inbuf[i] == NEWLINE { // as I use newline instead of ENDSTR, we have to ignore it
				continue
			}
			outbuf[j] = inbuf[i]
			j -= 1
			if j < 0 {
				j = MAXOUT - 1
			}
		}
		// copy thr part before FOLD to right
		j = MIDDLE + dist
		for i := 0; i < f; i++ {
			outbuf[j] = inbuf[i]
			j = j%(MAXOUT-1) + 1
		}
		last := 0
		for j = 0; j < MAXOUT; j++ { // find the last non-blank symbol to add ENDSTR after it
			if outbuf[j] != BLANK {
				last = j
			}
		}
		putstr(string(outbuf[0:last+1]), STDOUT) // We don't need ENDSTR
		putc(NEWLINE)
	}
}

func index(inbuf string, target uint8) int {
	n := len(inbuf)
	for i := 0; i < n; i++ {
		if inbuf[i] == target {
			return i
		}
	}
	return -1 // actually it should never happen because we except usage after kwic, it guarantees FOLD presence
}
