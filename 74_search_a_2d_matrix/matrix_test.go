package matrix

import "testing"

func TestSearchIt(t *testing.T) {
	st := []struct {
		name   string
		matrix [][]int
		target int
		exp    bool
	}{
		{"empty matrix", nil, 0, false},
		{"single element", [][]int{[]int{1}}, 2, false},
		{"testcase1", [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 60}}, 3, true},
		{"testcase2", [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 60}}, 13, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := searchMatrix(tt.matrix, tt.target)
			if out != tt.exp {
				t.Fatalf("with input matrix: %v and target:%d wanted %t but got %t", tt.matrix, tt.target, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
