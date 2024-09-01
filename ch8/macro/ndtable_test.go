package macro

import (
	"testing"
)

func Test_install_and_lookup(t *testing.T) {

	tests := []struct {
		token string
		defn  string
		kind  sttype
	}{
		{token: "define", defn: "", kind: DEFTYPE},
		{token: "xxx", defn: "yyy", kind: MACTYPE},
	}

	for _, test := range tests {
		inithash()
		var defn string
		var kind sttype

		install(test.token, test.defn, test.kind)
		result := lookup(test.token, &defn, &kind)
		if !result {
			t.Errorf("Nothing found for [%v, %v]", test.token, test.defn)
		}

		if test.defn != defn {
			t.Errorf("status: got %v want %v", defn, test.defn)
		}
		if test.kind != kind {
			t.Errorf("status: got %v want %v", kind, test.kind)
		}
	}

}
