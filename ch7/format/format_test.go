package format

import (
	"testing"
)

func Test_getcmd(t *testing.T) {
	tests := []struct {
		buf string
		cmd CmdType
	}{
		{buf: ".br", cmd: BR},
		{buf: ".break", cmd: BR},
		{buf: ".nf 123 123", cmd: NF},
		{buf: ".zz", cmd: UNKNOWN},
	}
	for _, test := range tests {
		result := getcmd(test.buf)

		if result != test.cmd {
			t.Errorf("status: got %v want %v", result, test.cmd)
		}
	}
}

func Test_gettl(t *testing.T) {
	tests := []struct {
		buf    string
		result string
	}{
		{buf: ".he Title", result: "Title"},
		{buf: ".header Title", result: "Title"},
		{buf: ".header '   Title", result: "   Title"},
		{buf: ".header \"   Title", result: "   Title"},

		{buf: ".fo Footer", result: "Footer"},
		{buf: ".footer Footer", result: "Footer"},
		{buf: ".fo '   Footer", result: "   Footer"},
		{buf: ".fo \"   Footer", result: "   Footer"},
	}
	for _, test := range tests {
		result := gettl(test.buf)

		if result != test.result {
			t.Errorf("status: got %v want %v", result, test.result)
		}
	}
}
