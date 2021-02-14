package matrix

import "testing"

func TestShortestPath(t *testing.T) {
	st := []struct {
		name   string
		matrix [][]int
		exp    int
	}{
		{"testcase1", [][]int{[]int{0, 1}, []int{1, 0}}, 2},
		{"testcase2", [][]int{[]int{0, 0, 0}, []int{1, 1, 0}, []int{1, 1, 0}}, 4},
		{"testcase3", [][]int{[]int{1, 0, 0}, []int{0, 0, 1}, []int{0, 0, 0}}, -1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := shortestPath(tt.matrix)
			if out != tt.exp {
				t.Fatalf("wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
