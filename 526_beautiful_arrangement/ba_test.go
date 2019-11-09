package ba

import "testing"

func TestBeautyArrangement(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"N=1", 1, 1},
		{"N=2", 2, 2},
		{"N=3", 3, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := beautifulArrangement(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}

}
