package main

const MAXCHARS = 10_000
const MAXLINES = 300

func inmemsort() {

	posbuf := make([]string, MAXLINES)
	res := gtext(posbuf, STDIN)
	if res {
		shell(posbuf)
		ptext(posbuf, STDOUT)
	}
}

func main() {

}
