package find

import (
	"testing"
)

// as it "main" package, run tests as > go test *.go
// go test makepat_test.go makepat.go  find.go io.go
func Test_dodash(t *testing.T) {
	tests := []struct {
		src      string
		delim    uint8
		expected string
	}{
		{src: "a-f]", delim: ']', expected: "abcdef"},
		{src: "abc-]", delim: ']', expected: "abc-"},
		{src: "a-dA-D]", delim: ']', expected: "abcdABCD"},
		{src: "@t]", delim: ']', expected: "\t"},
		{src: "@n]", delim: ']', expected: "\n"},
	}

	for _, test := range tests {
		result := dodash(test.src, test.delim)
		if result != test.expected {
			t.Errorf("got %v want %v", result, test.expected)
		}
	}
}

func Test_getccl(t *testing.T) {
	tests := []struct {
		arg      string
		i        int
		expected string
	}{
		//note, if [x____, x = 8, it its backspace, it was a problem to test
		{arg: "[a-f]", i: 0, expected: "[\x06abcdef"},
		{arg: "[^a-f]", i: 0, expected: "!\x06abcdef"},
		{arg: "[^a]", i: 0, expected: "!\x01a"},
	}

	for _, test := range tests {
		result := getccl(test.arg, test.i)
		if result != test.expected {
			t.Errorf("got %v want %v", result, test.expected)
		}
	}
}

func Test_stclose(t *testing.T) {
	tests := []struct {
		pat      string
		expected string
	}{
		{pat: "ca", expected: "*a"},
		{pat: "!\x01aca", expected: "!\x01a*a"},
	}

	for _, test := range tests {
		result := stclose(test.pat)
		if result != test.expected {
			t.Errorf("got %v want %v", result, test.expected)
		}
	}
}

func Test_Makepat(t *testing.T) {
	tests := []struct {
		arg      string
		start    int
		delim    uint8
		expected string
	}{
		{arg: "%|", start: 0, delim: '|', expected: "%"},
		{arg: "a|", start: 0, delim: '|', expected: "ca"},
		{arg: "*a|", start: 0, delim: '|', expected: "c*ca"},
		{arg: "$|", start: 0, delim: '|', expected: "$"},
		{arg: "?|", start: 0, delim: '|', expected: "?"},
		{arg: "$%|", start: 0, delim: '|', expected: "c$c%"},
		{arg: "%?a?$|", start: 0, delim: '|', expected: "%?ca?$"},
	}

	for _, test := range tests {
		result := Makepat(test.arg, test.start, test.delim)
		if result != test.expected {
			t.Errorf("got %v want %v", result, test.expected)
		}
	}
}
func Test_Makepat__ccl(t *testing.T) {
	tests := []struct {
		arg      string
		start    int
		delim    uint8
		expected string
	}{
		{arg: "a[^b]d|", start: 0, delim: '|', expected: "ca!\x01bcd"},
		{arg: "a[b]d|", start: 0, delim: '|', expected: "ca[\x01bcd"},
		{arg: "[^b]|", start: 0, delim: '|', expected: "!\x01b"},
	}

	for _, test := range tests {
		result := Makepat(test.arg, test.start, test.delim)
		if result != test.expected {
			t.Errorf("got %v want %v", result, test.expected)
		}
	}
}

func Test_Makepat__closure(t *testing.T) {
	tests := []struct {
		arg      string
		start    int
		delim    uint8
		expected string
	}{
		{arg: "a*|", start: 0, delim: '|', expected: "*a"},
		{arg: "a[b]d*c|", start: 0, delim: '|', expected: "ca[\x01b*dcc"},
		{arg: "a*b*c*|", start: 0, delim: '|', expected: "*a*b*c"},
	}

	for _, test := range tests {
		result := Makepat(test.arg, test.start, test.delim)
		if result != test.expected {
			t.Errorf("got %v want %v", result, test.expected)
		}
	}
}
