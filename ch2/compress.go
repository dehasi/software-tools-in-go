package main

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// putrep -- put out representation of run of n 'c's
func putrep(n int, c int8) {
	const MAXREP = 26
	const THRESH = 4

	for n >= THRESH || ((c == WARNING) && (n > 0)) {
		putc(WARNING)
		putc(int8(min(n, MAXREP) - 1 + 'A'))
		putc(c)
		n -= MAXREP
	}
	for ; n > 1; n-- {
		putc(c)
	}
}

// compress -- compresses standard input
func compress() {
	var c int8 = 0
	var lastc int8 = 0

	lastc = getc(&lastc)
	var n = 1
	for lastc != ENDFILE {
		if getc(&c) == ENDFILE {
			if (n > 1) || (lastc == WARNING) {
				putrep(n, lastc)
			} else {
				putc(lastc)
			}
		} else if c == lastc {
			n += 1
		} else if (n > 1) || (lastc == WARNING) {
			putrep(n, lastc)
			n = 1
		} else {
			putc(lastc)
		}
		lastc = c
	}
}
