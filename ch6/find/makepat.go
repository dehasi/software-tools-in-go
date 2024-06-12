package find

import (
	"ch6/io"
)

// makepat -- make pattern from arg[i], terminate at delim
func Makepat(arg string, start int, delim uint8) string {

	/*
	   junk : boolean;
	*/
	j := 0     // pat index,
	pat := ""  // pattern that we'll return
	i := start // arg index
	lastj := 0 // why not lastj = j?
	done := false
	for !done && arg[i] != delim && i < len(arg) {
		lj := j // maybe len(pat)
		if arg[i] == ANY {
			pat += string(ANY)
			j++
		} else if arg[i] == BOL && i == start {
			pat += string(BOL)
			j++
		} else if arg[i] == EOL && arg[i+1] == delim {
			pat += string(EOL)
			j++
		} else if arg[i] == CCL {
			cclPat := getccl(arg, i)
			j += len(cclPat) // maybe -1?
			pat += cclPat
			i = step_ccl(arg, i)
		} else if arg[i] == CLOSURE && i > start {
			lj = lastj
			if contains([]uint8{BOL, EOL, CLOSURE}, pat[lj]) {
				done = true // break?
			} else {
				// Important: we replace pat
				pat = stclose(pat)
			}
		} else {
			pat += string(LITCHAR)
			j++
			pat += string(esc(arg, i))
			j++
		}
		lastj = lj

		if !done {
			i = i + 1
		}
	}
	if done || arg[i] != delim { // finished early
		return "" // was -1
	} else {
		return pat // all is well
	}
}

// getccl -- expand char class at arg[i] into pat[j]
func getccl(arg string, i int) string {
	i = i + 1 // skip over ' ['
	result := ""
	if arg[i] == NEGATE {
		result += string(NCCL)
		i = i + 1
	} else {
		result += string(CCL)
	}
	expanded := dodash(arg[i:], CCLEND)
	result += string(uint8(len(expanded))) //  we expect that the length the expaneded string will fit into uint8 => 255 and in must because all ASCII is 127
	result += expanded
	return result
}

// step_ccl -- finds CCLEND after i
// we call it after getccl, to find when the process ended
// we need it to know to where shift index
func step_ccl(arg string, i int) int {
	if arg[i] == CCLEND {
		return i
	}
	for i < len(arg) {
		if arg[i] == CCLEND {
			return i
		}
		i++
	}
	io.Error("step_ccl: no CCLEND found")
	return i // never happen
}

// dodash - expands set at src[i] into dest[j], stop at delim
func dodash(src string, delim uint8) string {
	var result = ""
	for i := 0; src[i] != delim && i < len(src); i++ {
		if src[i] == ESCAPE {
			result += string(esc(src, i))
			i++ // we need to increment, if it's last it's ok. we end form the loop
		} else if src[i] != DASH {
			result += string(src[i])
		} else if i == len(src)-1 { // last element is just dash
			result += string(DASH)
		} else if src[i-1] < src[i+1] {
			for k := src[i-1] + 1; k < src[i+1]; k++ {
				result += string(k)
			}
		} else {
			result += string(DASH)
		}
	}

	return result
}

// esc -- maps s[i] into escaped character, increment i
func esc(s string, i int) uint8 {
	if s[i] != ESCAPE {
		return s[i]
	} else if i+1 == len(s) { // @ not special at end
		return ESCAPE
	} else {
		i++
		if s[i] == 'n' {
			return io.NEWLINE
		} else if s[i] == 't' {
			return io.TAB
		} else {
			return s[i]
		}
	}
}

// stclose -- insert closure entry at pat[j]
// input pat:[..,TYPE, VAL]
// expected input pat:[..,LITCHAR, a]
// expected output pat:[..,CLOSURE, a]
// we don't have groups, so, only one element for closure

func stclose(pat string) string {
	n := len(pat)
	return pat[0:n-2] + string(CLOSURE) + string(pat[n-1])
}

/*
p 52
addstr adds a character at a time to a specified position of an array and increments the index. It also checks that there's enough room to do so. We will use
this function extensively in later programs, so we will add it to our standard
context for all programs.

 addstr -- put c in outset[j] if it fits, increment j
function addstr(c : character; var outset: string;
    var j : integer; maxset : integer) : boolean;
begin
    if (j > maxset) then
        addstr := false
    else begin
        outset[j] : = c ;
        j := j + 1;
        addstr := true
    end
end;

Go Translation
// addstr -- put c in outset[j] if it fits, NOT increment j
func addstr(c uint8, outset *string, j int, maxset int) bool {
	var s string = *outset
	if j > len(s) || j > maxset {
		return false
	}

	x := s[:j] + string(c) + s[j+1:]
	*outset = x
	return true
}
*/
