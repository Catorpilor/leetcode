package marr

import "testing"

func TestMinArrEnd(t *testing.T) {
	st := []struct {
		name string
		n, x int
		exp  int
	}{
		{"single array", 1, 2, 2},
		{"testcase1", 3, 4, 6},
		{"testcase2", 2, 7, 15},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minArrEnd(tt.n, tt.x)
			if out != tt.exp {
				t.Fatalf("with input n: %d, x: %d wanted %d but got %d", tt.n, tt.x, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
