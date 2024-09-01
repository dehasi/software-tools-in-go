package macro

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_gettok(t *testing.T) {
	expected := []string{"define", "(", "x", ",", " ", "y", ")"}
	pbstr("define(x, y)")

	var tokens []string
	for token := gettok(100); token != ""; token = gettok(100) {
		tokens = append(tokens, token)
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("status: got %v want %v", strings.Join(tokens, ","), strings.Join(expected, ","))
	}
}

func Test_Macro_replace_in_out(t *testing.T) {
	input := lines(
		"define(sqr, $1 * $1)",
		"x = sqr(3)")
	expected := "x =  3 * 3"
	restoreStdin := replaceStdin(input)
	defer restoreStdin()

	output := captureStdout()

	Macro()

	actual := strings.TrimSpace(output())
	if actual != expected {
		t.Errorf("status: got %v want %v", actual, expected)
	}
}

func Test_Macro_simple_define(t *testing.T) {
	input := lines(
		"define(ENDFILE, (-1))",
		"define(DONE, ENDFILE)",
		"if (getit(line) = DONE) then",
		"    putit(sumline);")

	expected := lines(
		"if (getit(line) =   (-1)) then",
		"    putit(sumline);")

	restoreStdin := replaceStdin(input)
	defer restoreStdin()

	output := captureStdout()

	Macro()

	actual := strings.TrimSpace(output())
	if actual != expected {
		t.Errorf("status: got %v want %v", actual, expected)
	}
}

func Test_Macro_substr(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "x = substr(abc,2,1)", expected: "x = b"},
		{input: "x = substr(abc,2)", expected: "x = bc"},
		{input: "x = substr(abc,4)", expected: "x ="},
	}

	for _, test := range tests {
		restoreStdin := replaceStdin(test.input)

		defer restoreStdin()

		output := captureStdout()

		Macro()

		actual := strings.TrimSpace(output())
		if actual != test.expected {
			t.Errorf("status: got [%v] want [%v]", actual, test.expected)
		}
	}
}

func Test_Macro_ifelse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "x = ifelse(x,x,TRUE,FALSE)", expected: "x = TRUE"},
		{input: "x = ifelse(x,y,TRUE,FALSE)", expected: "x = FALSE"},
		{input: "x = ifelse(x,x,TRUE)", expected: "x = TRUE"},
		{input: "x = ifelse(x,y,TRUE)", expected: "x ="},
	}

	for _, test := range tests {
		restoreStdin := replaceStdin(test.input)

		defer restoreStdin()

		output := captureStdout()

		Macro()

		actual := strings.TrimSpace(output())
		if actual != test.expected {
			t.Errorf("status: got [%v] want [%v]", actual, test.expected)
		}
	}
}

func replaceStdin(input string) (restore func()) {
	oldStdin := os.Stdin
	rStdin, wStdin, _ := os.Pipe()
	os.Stdin = rStdin
	wStdin.WriteString(input)
	wStdin.Close() // simulate EOF

	return func() {
		os.Stdin = oldStdin
	}
}

func captureStdout() func() string {
	rStdout, wStdout, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = wStdout

	return func() string {
		os.Stdout = old
		wStdout.Close() // let the reader see EOF

		// Read the captured output
		var buf bytes.Buffer
		io.Copy(&buf, rStdout)
		return buf.String()
	}
}

func lines(parts ...string) string {
	return strings.Join(parts, "\n")
}
