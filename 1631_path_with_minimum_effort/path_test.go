package path

import "testing"

func TestMinimumEffort(t *testing.T) {
	st := []struct {
		name    string
		heights [][]int
		exp     int
	}{
		{"nil height", nil, 0},
		{"single row", [][]int{[]int{1, 2, 3}}, 1},
		{"testcase1", [][]int{[]int{1, 2, 2}, []int{3, 8, 2}, []int{5, 3, 5}}, 2},
		{"testcase2", [][]int{[]int{1, 2, 3}, []int{3, 8, 4}, []intP{5, 3, 5}}, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minimumEffort(tt.heights)
			if out != tt.exp {
				t.Fatalf("wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
