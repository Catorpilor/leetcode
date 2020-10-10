package comp

import "testing"

func TestBitwiseComplement(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"testcase1", 1, 0},
		{"testcase1", 5, 2},
		{"testcase3", 10, 5},
		{"testcase4", 7, 0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := bitwiseComp(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
