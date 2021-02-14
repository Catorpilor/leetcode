package matrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	st := []struct {
		name   string
		matrix [][]int
		target int
		exp    bool
	}{
		{"testcase1", [][]int{[]int{1, 4, 7, 11, 15}, []int{2, 5, 8, 12, 19}, []int{3, 6, 9, 16, 22}, []int{10, 13, 14, 17, 24}, []int{18, 21, 23, 26, 30}}, 5, true},
		{"testcase1", [][]int{[]int{1, 4, 7, 11, 15}, []int{2, 5, 8, 12, 19}, []int{3, 6, 9, 16, 22}, []int{10, 13, 14, 17, 24}, []int{18, 21, 23, 26, 30}}, 20, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := matrixSearch(tt.matrix, tt.target)
			if out != tt.exp {
				t.Fatalf("wanted %t but got %t", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
