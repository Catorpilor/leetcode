package cal

import "testing"

func TestCalculator(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  int
	}{
		{"testcase1", "1 + 1", 2},
		{"testcase2", "2 - 1+2", 3},
		{"testcase3", "(1+(4+5+2)-3)+(6+8)", 23},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := Calculator(tt.s)
			if out != tt.exp {
				t.Fatalf("with input s:%s wanted %d but got %d", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
