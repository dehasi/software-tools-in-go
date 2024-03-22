package main

// concat -- concatenates files onto standard output
func concat() {

	for i := 1; i < nargs(); i++ {
		s := getarg(i)
		fd := mustopenf(s)
		fcopy(fd, STDOUT)
	}
}
