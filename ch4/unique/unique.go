package main

const MAX_STR = 1000 // max symbols per line

// unique -- removes adjacent duplicate lines
func unique() {
	buf := make([]string, 2)
	cur := 0
	buf[1-cur] = "" // Go doesn't have '\0', and doesn't need
	got := false
	for {
		buf[cur], got = getline(STDIN, MAX_STR)
		if !got {
			return
		}
		if !equal(buf[cur], buf[1-cur]) {
			putstr(buf[cur], STDOUT)
		}
		cur = 1 - cur
	}
}

func equal(s1 string, s2 string) bool {
	return s1 == s2
}
func main() {
	unique()
}
