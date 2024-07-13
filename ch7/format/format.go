package format

import "ch7/io"

const CMD uint8 = '.'

func Format() {
	initfmt()
	for inbuf, result := io.Getline(io.STDIN, io.MAX_STR); result; inbuf, result = io.Getline(io.STDIN, io.MAX_STR) {
		if inbuf[0] == CMD {
			command(inbuf)
		} else {
			text(inbuf)
		}
	}
}

// command -- perform formatting command
func command(buf string) {
	panic("unimplemented")
}

func text(buf string) {
	io.Putstr(buf, io.STDOUT)
}

func initfmt() {

}
