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

func Test_docmd_subst(t *testing.T) {
	setbuf()
	defer clrbuf()
	tests := []struct {
		line string
		i    int
		glob bool
		// return
		status StCode

		// globals
		before globals
		after  globals
	}{
		// parses simple pattern
		{line: ".,.s/bbb/FFF/\n", i: 3, status: OK,
			before: globals{line1: 1, line2: 1, nlines: 2, curln: 1, lastln: 1, pat: ""},
			after:  globals{line1: 1, line2: 1, nlines: 2, curln: 1, lastln: 1, pat: "cbcbcb"}},
	}

	for _, test := range tests {
		puttxt("aaabbbccc\n")
		prepareState(test.before)

		status := docmd(test.line, test.i, test.glob)
		if status != test.status {
			t.Errorf("status: got %v want %v", status, test.status)
		}
		txt := gettxt(curln)
		if txt != "aaaFFFccc\n" {
			t.Errorf("status: got %v want %v", txt, "aaaFFFccc\n")
		}

		assert_gobals(t, test.after)
	}
}

func Test_docmd_subst_few_lines(t *testing.T) {
	setbuf()
	defer clrbuf()
	tests := []struct {
		line string
		i    int
		glob bool
		// return
		status StCode

		// globals
		before globals
		after  globals
	}{
		// parses simple pattern
		{line: "1,$s/bbb/FFF/g\n", i: 3, status: OK,
			before: globals{line1: 1, line2: 5, nlines: 5, curln: 5, lastln: 5, pat: ""},
			after:  globals{line1: 1, line2: 5, nlines: 5, curln: 5, lastln: 5, pat: "cbcbcb"}},
	}

	for _, test := range tests {
		puttxt("bbbfffddd\n")
		puttxt("ababcc\n")
		puttxt("ccfffaaa\n")
		puttxt("aaabbbccc\n")
		puttxt("bbbbbbbbb\n")
		prepareState(test.before)

		status := docmd(test.line, test.i, test.glob)
		if status != test.status {
			t.Errorf("status: got %v want %v", status, test.status)
		}

		assert_equals(gettxt(1), "FFFfffddd\n", t)
		assert_equals(gettxt(2), "ababcc\n", t)
		assert_equals(gettxt(3), "ccfffaaa\n", t)
		assert_equals(gettxt(4), "aaaFFFccc\n", t)
		assert_equals(gettxt(5), "FFFFFFFFF\n", t)

		assert_gobals(t, test.after)
	}
}

func Test_getnum_contextSerch(t *testing.T) {
	setbuf()
	defer clrbuf()
	tests := []struct {
		line string
		i    int
		// return
		num    int
		ii     int
		status StCode
		// globals
		// global state
		before globals
		after  globals
	}{
		// parses simple pattern
		{line: "\\bb\\p", i: 0, status: OK, ii: 4, num: 2,
			before: globals{line1: -1, line2: -1, nlines: 0, curln: 0, lastln: 0, pat: ""},
			after:  globals{line1: -1, line2: -1, nlines: 0, curln: 4, lastln: 4, pat: "cbcb"}},
	}

	for _, test := range tests {
		prepareState(test.before)
		puttxt("1aaa\n")
		puttxt("2bbb\n")
		puttxt("3ccc\n")
		puttxt("4ddd\n")

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

		assert_gobals(t, test.after)
	}
}

func Test_getnum(t *testing.T) {
	setbuf()
	defer clrbuf()
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
		// no numbers returns ENDDATA
		{line: "p", i: 0, status: ENDDATA, num: 0, ii: 0, line1: 1, line2: 2, nlines: 4, curln: 3, lastln: 4},
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
	setbuf()
	defer clrbuf()
	tests := []struct {
		line string
		i    int
		// return
		ii     int
		status StCode
		// global state
		before globals
		after  globals
	}{
		// reads curln
		{line: ".p", i: 0, status: OK, ii: 1,
			before: globals{line1: -1, line2: -1, nlines: 505, curln: 101, lastln: 504, pat: ""},
			after:  globals{line1: 101, line2: 101, nlines: 1, curln: 101, lastln: 504, pat: ""}},
		// reads curln +42
		{line: ".+42p", i: 0, status: OK, ii: 4,
			before: globals{line1: -1, line2: -1, nlines: 505, curln: 101, lastln: 504, pat: ""},
			after:  globals{line1: 143, line2: 143, nlines: 1, curln: 101, lastln: 504, pat: ""}},
		// reads curln +27
		{line: ".-27p", i: 0, status: OK, ii: 4,
			before: globals{line1: -1, line2: -1, nlines: 505, curln: 101, lastln: 504, pat: ""},
			after:  globals{line1: 74, line2: 74, nlines: 1, curln: 101, lastln: 504, pat: ""}},

		// parses comma, reads line1 = curln, line2 = lastln
		{line: ".,$p", i: 0, status: OK, ii: 3,
			before: globals{line1: -1, line2: -1, nlines: 505, curln: 101, lastln: 504, pat: ""},
			after:  globals{line1: 101, line2: 504, nlines: 2, curln: 101, lastln: 504, pat: ""}},

		// parses comma,
		{line: ".+42,$-24p", i: 0, status: OK, ii: 9,
			before: globals{line1: -1, line2: -1, nlines: 505, curln: 101, lastln: 504, pat: ""},
			after:  globals{line1: 101 + 42, line2: 504 - 24, nlines: 2, curln: 101, lastln: 504, pat: ""}},
	}

	for _, test := range tests {
		prepareState(test.before)

		i, status := getlist(test.line, test.i)
		if status != test.status {
			t.Errorf("status: got %v want %v", status, test.status)
		}
		if i != test.ii {
			t.Errorf("i: got %v want %v", i, test.ii)
		}
		assert_gobals(t, test.after)
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

func assert_equals(actual string, extected string, t *testing.T) {
	if actual != extected {
		t.Errorf("status: got %v want %v", actual, extected)
	}
}
