package main

// expand -- uncompresses  standard input
func expand() {
	var c int8 = 0

	for getc(&c) != ENDFILE {
		if c != WARNING {
			putc(c)
		} else if is_uppper(getc(&c)) {
			var n = int(c - 'A' + 1)
			if getc(&c) != ENDFILE {
				for ; n > 1; n-- {
					putc(c)
				}
			} else {
				putc(WARNING)
				putc(int8(n - 1 + 'A'))
			}
		}
	}
}

func is_uppper(c int8) bool {
	return 'A' <= c && c <= 'Z'
}
