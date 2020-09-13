package pos

import "testing"

func TestSpecialPos(t *testing.T) {
	st := []struct {
		name string
		mat  [][]int
		exp  int
	}{
		{"single", [][]int{[]int{0}}, 0},
		{"testcase1", [][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{0, 0, 1}}, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := specialPos(tt.mat)
			if out != tt.exp {
				t.Fatalf("with input mat:%v wanted %d but got %d", tt.mat, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
