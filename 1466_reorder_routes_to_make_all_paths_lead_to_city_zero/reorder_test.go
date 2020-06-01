package reorder

import "testing"

func TestMinReorder(t *testing.T) {
	st := []struct {
		name  string
		n     int
		conns [][]int
		exp   int
	}{
		{"testcase1", 2, [][]int{[]int{1, 0}}, 0},
		{"testcase2", 2, [][]int{[]int{0, 1}}, 1},
		{"testcase3", 6, [][]int{[]int{0, 1}, []int{1, 3}, []int{2, 3}, []int{4, 0}, []int{4, 5}}, 3},
		{"testcase4", 5, [][]int{[]int{1, 0}, []int{1, 2}, []int{3, 2}, []int{3, 4}}, 2},
		{"testcase5", 3, [][]int{[]int{1, 0}, []int{2, 0}}, 0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minReorder(tt.n, tt.conns)
			if out != tt.exp {
				t.Fatalf("with input n:%d and conns: %v wanted %d but got %d", tt.n, tt.conns, tt.exp, out)
			}
			t.Log("Pass")
		})
	}
}
