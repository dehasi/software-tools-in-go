package main

import (
	"testing"
)

// as it "main" package, run tests as > go test *.go
func Test_match(t *testing.T) {
	tests := []struct {
		line     string
		pattern  string
		expected bool
	}{
		{line: "aaaa", pattern: "ca", expected: true},
		{line: "bbbb", pattern: "ca", expected: false},
		{line: "bAbb", pattern: "cA", expected: true},
		{line: "abcd", pattern: "ca[\x02ab", expected: true},
		{line: "abcd", pattern: "ca!\x02ac", expected: true},
		{line: "abcd", pattern: "ca?cc", expected: true},
		{line: "abcd", pattern: "%ca", expected: true},
		{line: "abcd", pattern: "%cb", expected: false},
		{line: "abcd", pattern: "ca$", expected: false},
		{line: "abcd", pattern: "cd$", expected: true},
	}

	for _, test := range tests {
		result := match(test.line, test.pattern)
		if result != test.expected {
			t.Errorf("got %v want %v", result, test.expected)
		}
	}
}

func Test_match_closure(t *testing.T) {
	tests := []struct {
		line     string
		pattern  string
		expected bool
	}{
		{line: "b", pattern: "*a", expected: true},
		{line: "ba", pattern: "*a", expected: true},
		{line: "bab", pattern: "*a", expected: true},
		{line: "ba", pattern: "*a*b", expected: true},
		{line: "bab", pattern: "*a*bcb", expected: true},
		{line: "bacde", pattern: "*a*b", expected: true},
	}

	for _, test := range tests {
		result := match(test.line, test.pattern)
		if result != test.expected {
			t.Errorf("got %v want %v", result, test.expected)
		}
	}
}

func Test_amatch(t *testing.T) {
	tests := []struct {
		line     string
		i        int
		pattern  string
		j        int
		expected int
	}{
		{line: "aaaa", i: 0, pattern: "ca", j: 0, expected: 1},
		{line: "bbbb", i: 0, pattern: "ca", j: 0, expected: -1},
		{line: "bAbb", i: 1, pattern: "cA", j: 0, expected: 2},
	}

	for _, test := range tests {
		result := amatch(test.line, test.i, test.pattern, test.j)
		if result != test.expected {
			t.Errorf("got %v want %v", result, test.expected)
		}
	}
}
