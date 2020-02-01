package lipm

import "testing"

func TestLongestIncPath(t *testing.T) {
	st := []struct {
		name   string
		matrix [][]int
		exp    int
	}{
		{"empty matrix", [][]int{}, 0},
		{"testcase1", [][]int{[]int{9, 9, 4}, []int{6, 6, 8}, []int{2, 1, 1}}, 4},
		{"testcase2", [][]int{[]int{3, 4, 5}, []int{3, 2, 6}, []int{2, 2, 1}}, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := longestIncPath(tt.matrix)
			if out != tt.exp {
				t.Fatalf("with matrix: %v wanted %d but got %d", tt.matrix, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
