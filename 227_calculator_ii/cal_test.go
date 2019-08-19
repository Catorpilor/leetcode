package cal

import "testing"

func TestCalculate(t *testing.T) {
	st := []struct {
		name  string
		input string
		exp   int
	}{
		{"empty string", "", 0},
		{"testcase1", "3+2*2", 7},
		{"testcase2", " 3/2 ", 1},
		{"testcase3", "3+5   /  2", 5},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := Calculate(tt.input)
			if out != tt.exp {
				t.Fatalf("with input %s wanted %d but got %d", tt.input, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
