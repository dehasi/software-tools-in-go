package main

// makecopy - copies one file to another
func makecopy() {
	if nargs() != 3 {
		error("usage: makecopy old new")
		return
	}
	inname := getarg(1)
	outname := getarg(2)

	if inname == outname {
		return
	}

	fin := mustopenf(inname)
	fout := mustcreatef(outname)
	fcopy(fin, fout)

	fin.Close()
	fout.Close()
}
