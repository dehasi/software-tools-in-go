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

func Test_spread(t *testing.T) {
	tests := []struct {
		buf    string
		nextra int
		result string

		outwds int
		outp   int
	}{
		{buf: "a b c", outwds: 3, outp: 5, nextra: 2, result: "a  b  c"},
	}
	for _, test := range tests {
		outwds = test.outwds
		outp = test.outp

		result := spread(test.buf, test.nextra)

		if result != test.result {
			t.Errorf("status: got [%v] want [%v]", result, test.result)
		}
	}
}
