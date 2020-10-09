package intervals

import "testing"

func TestCoverdIntervals(t *testing.T) {
	st := []struct {
		name      string
		intervals [][]int
		exp       int
	}{
		{"testcase1", [][]int{[]int{1, 4}, []int{3, 6}, []int{2, 8}}, 2},
		{"testcase2", [][]int{[]int{1, 4}, []int{2, 3}}, 1},
		{"testcase3", [][]int{[]int{0, 10}, []int{5, 12}}, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := coveredIntervals(tt.intervals)
			if out != tt.exp {
				t.Fatalf("with input intervals: %v wanted %d but got %d", tt.intervals, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
