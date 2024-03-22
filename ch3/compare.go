package main

// compare (simple version) -- compares two files for equality
func compare() {
	arg1 := getarg(1)
	arg2 := getarg(2)
	infile1 := mustopenb(arg1)
	infile2 := mustopenb(arg2)

	var lineno = 0
	for {
		lineno += 1
		line1, f1 := getline(infile1, MAXLINE)
		line2, f2 := getline(infile2, MAXLINE)

		if f1 && f2 {
			if line1 != line2 {
				diffmsg(lineno, line1, line2)
			}
		}
		if !f1 || !f2 {
			break
		}
		if f2 && !f1 {
			message("compare: end if file1")
		} else if f1 && !f2 {
			message("compare: end if file2")
		}
	}

}

// diffmsg -- prints line numbers and differing lines
func diffmsg(n int, line1 string, line2 string) {
	putdec(n, 1)
	putc(COLON)
	putc(NEWLINE)
	putstr(line1, STDOUT)
	putc(NEWLINE)
	putstr(line2, STDOUT)
	putc(NEWLINE)
}
