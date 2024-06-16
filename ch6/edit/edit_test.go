package edit

import (
	"testing"
)

func prepareState(test globals) {
	line1 = test.line1
	line2 = test.line2
	nlines = test.nlines
	curln = test.curln
	lastln = test.lastln
	pat = test.pat
}

func Test_getnum(t *testing.T) {
	tests := []struct {
		line string
		i    int
		// return
		num    int
		ii     int
		status StCode
		// globals
		line1  int // first line number
		line2  int // second line number
		nlines int // # of line numbers specified
		curln  int // current line -- value of dot
		lastln int
	}{
		// returns curln as num, increments i
		{line: ".p", i: 0, status: OK, num: 3, ii: 1, line1: 1, line2: 2, nlines: 4, curln: 3, lastln: 4},
		// returns lastln as num, increments i
		{line: "$p", i: 0, status: OK, num: 4, ii: 1, line1: 1, line2: 2, nlines: 4, curln: 3, lastln: 4},
		// parses a number into num
		{line: "42p", i: 0, status: OK, num: 42, ii: 2, line1: 1, line2: 2, nlines: 4, curln: 3, lastln: 4},
	}

	for _, test := range tests {
		prepareState(globals{
			line1:  test.line1,
			line2:  test.line2,
			nlines: test.nlines,
			curln:  test.curln,
			lastln: test.lastln,
			pat:    "",
		})
		num, i, status := getnum(test.line, test.i)
		if status != test.status {
			t.Errorf("status: got %v want %v", status, test.status)
		}
		if i != test.ii {
			t.Errorf("i: got %v want %v", i, test.ii)
		}
		if num != test.num {
			t.Errorf("num: got %v want %v", num, test.num)
		}

		assert_gobals(t, globals{
			line1:  test.line1,
			line2:  test.line2,
			nlines: test.nlines,
			curln:  test.curln,
			lastln: test.lastln,
			pat:    "",
		})
	}
}
func Test_getlist(t *testing.T) {
	tests := []struct {
		line string
		i    int
		// return
		ii     int
		status StCode
		line1  int // first line number
		line2  int // second line number
		nlines int // # of line numbers specified
		curln  int // current line -- value of dot
		lastln int
	}{
		{line: ".p", i: 0, status: OK, line1: 0, line2: 0, nlines: 0, curln: 1, lastln: 0},
	}

	for _, test := range tests {
		i, status := getlist(test.line, test.i)
		if status != test.status {
			t.Errorf("status: got %v want %v", status, test.status)
		}
		if i != test.ii {
			t.Errorf("i: got %v want %v", i, test.ii)
		}
		assert_gobals(t, globals{
			line1:  test.line1,
			line2:  test.line2,
			nlines: test.nlines,
			curln:  test.curln,
			lastln: test.lastln,
			pat:    "",
		})
	}
}

type globals struct {
	line1  int
	line2  int
	nlines int
	curln  int
	lastln int
	pat    string
}

func assert_gobals(t *testing.T, test globals) {
	if line1 != test.line1 {
		t.Errorf("line1: got %v want %v", line1, test.line1)
	}

	if line2 != test.line2 {
		t.Errorf("line2: got %v want %v", line2, test.line2)
	}

	if nlines != test.nlines {
		t.Errorf("nlines: got %v want %v", nlines, test.nlines)
	}

	if curln != test.curln {
		t.Errorf("curln: got %v want %v", curln, test.curln)
	}

	if lastln != test.lastln {
		t.Errorf("lastln: got %v want %v", lastln, test.lastln)
	}

	if pat != test.pat {
		t.Errorf("pat: got %v want %v", pat, test.pat)
	}
}
