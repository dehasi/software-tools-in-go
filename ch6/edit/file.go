package edit

import (
	"ch6/io"
)

// getfn -- get file name from line[i]
func getfn(line string, i int) (string, StCode) {
	if line[i+1] == io.BLANK {
		file, k := io.Getword(line, i+2) // get new filename
		if k > 0 && line[k] == io.NEWLINE {
			if len(savefile) <= 0 {
				savefile = file
			}
			return savefile, OK
		}
	} else if line[i+1] == io.NEWLINE && len(savefile) > 0 {
		return savefile, OK
	}
	return "", ERR
}

// doread -- read "file" after line n
func doread(n int, file string) StCode {
	return ERR
}

// dowrite -- write lines n1 .. n2 into file
func dowrite(n1 int, n2 int, file string) StCode {
	return ERR
}
