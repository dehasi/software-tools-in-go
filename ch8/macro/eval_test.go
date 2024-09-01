package macro

import (
	"testing"
)

func Test_expr(t *testing.T) {
	tests := []struct {
		expr   string
		result int
	}{
		{expr: "2+2*2", result: 6},
		{expr: "(2+2)*2", result: 8},
	}

	for _, test := range tests {
		i := 0
		result := expr(test.expr, &i)
		if result != test.result {
			t.Errorf("In %v: got %v want %v", test.expr, test.result, result)
		}
	}

}
