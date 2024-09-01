package macro

import "ch8/io"

const HASHSIZE = 53 // size of hash table
// deftype -- type definitions for define
type ndptr = *ndblock
type ndblock struct {
	name    charpos
	defn    charpos
	kind    sttype
	nextptr ndptr
}

// defvar -- var declarations for define
// TODO: consider

var hashtab [HASHSIZE]ndptr
var ndtable charbuf
var nexttab charpos // first free position in ndtable

func inithash() {
	nexttab = 1
	for i := 1; i < HASHSIZE; i++ {
		hashtab[i] = nil
	}
}

func lookup(name string, defn *string, kind *sttype) bool {
	p := hashfind(name)
	if p == nil {
		return false
	}
	*defn = cscopy(ndtable[:], p.defn)
	*kind = p.kind
	return true
}

// install -- add name, definition and type to table
func install(name, defn string, t sttype) {
	defn = withEndStr(defn)
	nlen := len(name) + 1
	dlen := len(defn) + 1

	if nexttab+nlen+dlen > MAXCHARS {
		io.Putstr(name, io.STDERR)
		io.Error(": too many definitions")
	}
	// put it at front of chain
	h := hash(name)
	p := new(ndblock)
	p.nextptr = hashtab[h]
	hashtab[h] = p
	p.name = nexttab
	sccopy(ndtable[:], name, nexttab)
	nexttab += nlen
	p.defn = nexttab
	sccopy(ndtable[:], defn, nexttab)
	nexttab += dlen
	p.kind = t
}

func withEndStr(defn string) string {
	if len(defn) == 0 || defn[len(defn)-1] != '\n' {
		return defn + "\n"
	}
	return defn
}

// hashfind -- find name in hash table
func hashfind(name string) *ndblock {
	for p := hashtab[hash(name)]; p != nil; p = p.nextptr {
		var tempname string = cscopy(ndtable[:], p.name)
		if tempname == name {
			return p
		}
	}
	return nil

}

// { hash -- compute hash function of a name
func hash(name string) int {
	h := 0
	for i := 0; i < len(name); i++ {
		h = (2*h + int(name[i])) % HASHSIZE
	}
	return h
}
