package gen

import "testing"

func TestGetMaximumGenerated(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"n=0", 0, 0},
		{"n=1", 1, 1},
		{"testcase1", 7, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := getMaximumGen(tt.n)
			if out != tt.exp {
				t.Fatalf("wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
