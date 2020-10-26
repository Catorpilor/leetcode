package domino

import "testing"

func TestMinRotations(t *testing.T) {
	st := []struct {
		name string
		a, b []int
		exp  int
	}{
		{"testcase1", []int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}, 2},
		{"testcase2", []int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}, -1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minRotations(tt.a, tt.b)
			if out != tt.exp {
				t.Fatalf("with input piles a:%v b:%v wanted %d but got %d", tt.a, tt.b, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
