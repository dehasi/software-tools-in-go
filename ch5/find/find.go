package main

const MAXPAT = MAX_STR
const CLOSIZE = 1
const CLOSURE uint8 = '*'
const BOL uint8 = '%'
const EOL uint8 = '$'
const ANY uint8 = '?'
const CCL uint8 = '['
const CCLEND uint8 = ']'
const NEGATE uint8 = '^'
const NCCL uint8 = '!'
const LITCHAR uint8 = 'c'

// dodash
const ESCAPE uint8 = '@'
const DASH uint8 = '-'

// makepat
const ENDSTR uint8 = '\n'

// change -- change "from" into "to" on each line }
func change() {
	if nargs() != 2 && nargs() != 3 {
		error("usage: change from [to]")
	}
	pat, ok := getpat(getarg(1))
	if !ok {
		error("change: illegal 'from' pattern")
	}
	sub := ""
	if nargs() == 3 {
		sub, ok = getsub(getarg(2))
		if !ok {
			error("change: illegal 'to' pattern")
		}
	}

	for line, lineRead := getline(STDIN, MAX_STR); lineRead; line, lineRead = getline(STDIN, MAX_STR) {
		subline(line, pat, sub)
	}
}

// find -- finds patterns in text
func find() {
	if nargs() != 2 {
		error("usage: find pattern")
	}
	pattern, ok := getpat(getarg(1))
	if !ok {
		error("find: illegal pattern")
	}

	for line, lineRead := getline(STDIN, MAX_STR); lineRead; line, lineRead = getline(STDIN, MAX_STR) {
		if match(line, pattern) {
			putstr(line, STDOUT)
		}
	}
}

// match -- finds match anywhere on the line
func match(line string, pattern string) bool {
	pos := -1
	n := len(line)
	for i := 0; i < n && pos == -1; i++ {
		pos = amatch(line, i, pattern, 0)
	}
	return pos >= 0
}

// amatch -- looks for match of pat[j] ... at line[offset] .
func amatch(line string, offset int, pattern string, j int) int {
	n := len(pattern)
	for j < n {
		if pattern[j] == CLOSURE {
			// j = j + patsize(pattern, j) // step over CLOSURE, why do we need it?
			i := offset
			// match as many as possible
			for i < len(line) {
				match, newI := omatch(line, i, pattern, j)
				i = newI
				if !match {
					break
				}
			}
			// i points to input character that made us fail
			// match rest of pattern against rest of input
			// shrink closure by 1 after each failure }
			k := -1
			for i >= offset {
				k = amatch(line, i, pattern, j+patsize(pattern, j))
				if k >= 0 { // matched rest of pattern
					break
				} else {
					i = i - 1
				}
			}
			return k
		} else {
			match, newOffset := omatch(line, offset, pattern, j)
			offset = newOffset
			if !match {
				offset = -1 // non-closure
				return -1
			} else { // omatch succeeded on this pattern element
				j += patsize(pattern, j)
			}
		}
	}
	return offset
}

// patsize -- returns size of pattern entry at pattern[n]
func patsize(pattern string, n int) int {
	switch pattern[n] {
	case LITCHAR:
		return 2
	case BOL, EOL, ANY:
		return 1
	case CCL, NCCL:
		return int(pattern[n+1] + 2) // [,count,a,b,c -> that's why +2
	case CLOSURE:
		return 1 + CLOSIZE // closure is the same as LITCHAR
	default:
		error("in patsize: can't happen: " + string(pattern[n]))
		return 42
	}
}

// omatch -- match one pattern element at pat[j]
func omatch(line string, i int, pattern string, j int) (bool, int) {

	line_size := len(line)
	if i > line_size {
		return false, i
	} else if !contains([]uint8{LITCHAR, BOL, EOL, ANY, CCL, NCCL, CLOSURE}, pattern[j]) {
		error("in omatch: can't happen: " + string(pattern[j]))
		return false, i
	}

	advance := -1
	switch pattern[j] {
	case LITCHAR, CLOSURE: // Added CLOSURE myself
		if line[i] == pattern[j+1] {
			advance = 1
		}
	case BOL:
		if i == 0 {
			advance = 0
		}
	case ANY:
		if i < line_size {
			advance = 1
		}
	case EOL:
		if i == line_size {
			advance = 0
		}
	case CCL:
		if locate(line[i], pattern, j+1) {
			advance = 1
		}
	case NCCL:
		if i < line_size && !locate(line[i], pattern, j+1) {
			advance = 1
		}
	default:
		error("in omatch: can't happen")
	}

	if advance >= 0 {
		i += advance
		return true, i
	}
	return false, i
}

// locate -- look for c in character class at pat[offset]
func locate(c uint8, pattern string, offset int) bool {
	for i := offset + int(pattern[offset]); i > offset; i-- {
		if c == pattern[i] {
			return true
		}
	}
	return false
}

func contains(array []uint8, val uint8) bool {
	for _, elem := range array {
		if elem == val {
			return true
		}
	}
	return false
}
func getpat(arg string) (string, bool) {
	return makepat(arg+"\n", 0, NEWLINE), true
}

func main() {
	// find()
	change()
}
