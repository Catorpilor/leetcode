package validp

import "testing"

func TestIsValid(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  bool
	}{
		{"empty string", "", true},
		{"single character", "(", false},
		{"testcase1", "()[]", true},
		{"testcase2", "({[]})", true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isValid(tt.s)
			if out != tt.exp {
				t.Fatalf("with input s:%s wanted %t but got %t", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
