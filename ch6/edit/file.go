package edit

import (
	"ch6/io"
	"os"
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
	fd, err := os.Open(file)
	if err != nil {
		return ERR
	}
	count := 0
	for inline, t := io.Getline(fd, io.MAXLINE); t; inline, t = io.Getline(fd, io.MAXLINE) {
		stat := puttxt(inline)
		if stat != OK {
			return stat
		}
		count++
	}
	io.Close(fd)
	io.Putdec(count, 1)
	io.Putc(io.NEWLINE)
	return OK
}

// dowrite -- write lines n1 .. n2 into file
// replaces if file not empty
func dowrite(n1 int, n2 int, file string) StCode {
	fd, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		println("err:", err)
		return ERR
	}

	for i := n1; i <= n2; i++ {
		line := gettxt(i)
		io.Putstr(line, fd)
	}
	io.Close(fd)
	io.Putdec(n2-n1+1, 1)
	io.Putc(io.NEWLINE)
	return OK
}
