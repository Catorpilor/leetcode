package noi

import "testing"

func TestMinErase(t *testing.T) {
	st := []struct {
		name      string
		intervals [][]int
		exp       int
	}{
		{"nil intervals", nil, 0},
		{"single one", [][]int{[]int{1, 2}}, 0},
		{"testcase1", [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{1, 3}}, 1},
		{"testcase2", [][]int{[]int{1, 2}, []int{1, 2}, []int{1, 2}}, 2},
		{"testcase3", [][]int{[]int{1, 2}, []int{2, 3}}, 0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minErase(tt.intervals)
			if out != tt.exp {
				t.Fatalf("with input intervals: %v wanted %d but got %d", tt.intervals, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
