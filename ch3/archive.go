package main

import (
	"os"
)

const MAXFILES = 100

var fname = make([]string, 1)
var fstat = make([]bool, 1)
var nfiles = 0
var errcount = 0
var archtemp = "atemp"
var archhdr = "-h-"

// - c create a new archive with named members
// - d delete named members from archive
// - p print named members on standard output
// - t print table of archive contents
// - u update named members or add at end
// - x extract named members from archive
func archive() {
	if nargs() < 3 {
		help()
	}
	cmd := getarg(1)
	aname := getarg(2)

	getfns()
	if len(cmd) != 2 || cmd[0] != '-' {
		help()
	} else if cmd[1] == 'c' || cmd[1] == 'u' {
		update(aname, cmd[1])
	} else if cmd[1] == 't' {
		table(aname)
	} else if cmd[1] == 'x' || cmd[1] == 'p' {
		extract(aname, cmd[1])
	} else if cmd[1] == 'd' {
		delete_archive(aname)
	} else {
		help()
	}
}

func getfns() {
	nfiles = nargs() - 3
	if nfiles > MAXFILES {
		error("archive: to many file names")
	}
	fname = make([]string, nfiles)
	fstat = make([]bool, nfiles)
	for i := 0; i < nfiles; i++ {
		fname[i] = getarg(i + 3)
		fstat[i] = false
	}
	for i := 0; i < nfiles-1; i++ {
		for j := i + 1; j < nfiles; j++ {
			if fname[i] == fname[j] {
				putstr(fname[i], STDERR)
				error(": duplicate file name")
			}
		}
	}

}

// update -- updates existing files, add new ones at the end
func update(aname string, cmd uint8) {
	tfd := mustcreatef(archtemp)
	if cmd == 'u' {
		afd := mustopenf(aname)
		replace(afd, tfd, 'u') // update existing
		close(afd)
	}
	for i := 0; i < nfiles; i++ {
		if !fstat[i] {
			addfile(fname[i], tfd)
			fstat[i] = true
		}
	}
	close(tfd)
	if errcount == 0 {
		fmove(archtemp, aname)
	} else {
		message("fatal errors  - archive not altered")
	}
	remove(archtemp)
}

// fmove -- move file name1 to name2
func fmove(name1 string, name2 string) {
	fd1 := mustopenf(name1)
	fd2 := mustcreatef(name2)
	fcopy(fd1, fd2)
	close(fd1)
	close(fd2)
}

// addfile -- adds file "name" to archive
func addfile(name string, fd *os.File) {
	nfd, err := os.Open(name)
	if err != nil {
		putstr(name, STDERR)
		message(": can't add")
		errcount += 1
	}
	if errcount == 0 {
		var head = makehdr(name)
		putstr(head, fd)
		fcopy(nfd, fd)
		close(nfd)
	}
}

func makehdr(name string) string {
	const blank = " "
	return archhdr + blank +
		name + blank +
		itoc(fsize(name)) + "\n"
}

func fsize(name string) int {
	var result = 0
	var c int8 = ' '
	fd := mustopenf(name)
	for getcf(&c, fd) != ENDFILE {
		result += 1
	}
	close(fd)
	return result
}

// replace -- replaces or delete files }
func replace(afd *os.File, tfd *os.File, cmd uint8) {

	for {
		inline, uname, size, isheader := gethdr(afd)
		if !isheader {
			break
		}
		if filearg(uname) {
			if cmd == 'u' { // add new one
				addfile(uname, tfd)
			}
			fskip(afd, size) // discard old file
		} else {
			putstr(inline, tfd)
			putcf(NEWLINE, tfd)
			acopy(afd, tfd, size)
		}
	}
}

// table -- prints table of archive contents
func table(aname string) {
	afd := mustopenf(aname)
	for {
		head, name, size, isheader := gethdr(afd)
		if !isheader {
			break
		}
		if filearg(name) {
			tprint(head)
		}
		fskip(afd, size)
	}
	notfound()
}

// notfound -- prints "not found" warning
func notfound() {
	for i := 0; i < nfiles; i++ {
		if !fstat[i] {
			putstr(fname[i], STDERR)
			message(": not in archive")
			errcount += errcount
		}
	}
}

// filearg -- checks if name matches argument list
func filearg(name string) bool {
	if nfiles <= 0 {
		return true
	}
	for i := 0; i < nfiles; i++ {
		if fname[i] == name {
			fstat[i] = true
			return true
		}
	}
	return false
}

func fskip(afd *os.File, size int) {
	var c int8 = 0
	for i := 0; i < size; i++ {
		if getcf(&c, afd) == ENDFILE {
			error("archive: end of file in skip")
		}
	}
}

func gethdr(fd *os.File) (string, string, int, bool) {
	buf, done := getlinef(fd, MAXLINE)
	if done == false {
		return "", "", -1, false
	}
	tmp, i := getword(buf, 0)
	if tmp != archhdr {
		error("archive is not in proper format, tmp:" + tmp)
	}
	name, i := getword(buf, i)
	size, i := getword(buf, i)
	return buf, name, ctoi(size), true
}

// tprint -- prints table entry for one member
func tprint(buf string) {
	_, i := getword(buf, 0) // skip -hdr-
	name, i := getword(buf, i)
	putstr(name, STDOUT)
	putc(BLANK)
	size, i := getword(buf, i)
	putstr(size, STDOUT)
	putc(NEWLINE)
}

// extract -- extracts files from archive
func extract(aname string, cmd uint8) {
	afd := mustopenf(aname)
	efd := STDERR
	if cmd == 'p' {
		efd = STDOUT
	}

	for {
		_, ename, size, isheader := gethdr(afd)
		if !isheader {
			break
		}
		if !filearg(ename) {
			fskip(afd, size)
		} else {
			if efd != STDOUT {
				efd = create(ename)
			}
			if efd == nil {
				putstr(ename, STDERR)
				message("Can't create")
				errcount += 1
				fskip(afd, size)
			} else {
				acopy(afd, efd, size)
				if efd != STDOUT {
					close(efd)
				}
			}
		}
	}
	notfound()
}

func acopy(fdi *os.File, fdo *os.File, size int) {
	var c int8 = 0
	for i := 0; i < size; i++ {
		if getcf(&c, fdi) == ENDFILE && i != size-1 {
			error("archive: end of file in acopy" + " i = " + itoc(i))
		}
		putcf(c, fdo)
	}
}

// 'delete' is a reserved word for go
func delete_archive(aname string) {
	if nfiles <= 0 {
		error("archive: -d requires explicit file names")
	}
	afd := mustopenf(aname)
	tfd := mustcreatef(archtemp)

	replace(afd, tfd, 'd')

	close(afd)
	close(tfd)

	if errcount == 0 {
		fmove(archtemp, aname)
	} else {
		message("fatal errors - archivie not altered")
	}
	remove(archtemp)

}

// help -- prints diagnostic for archive
func help() {
	error("usage: archive -[cdptux] archname [files....]")
}
