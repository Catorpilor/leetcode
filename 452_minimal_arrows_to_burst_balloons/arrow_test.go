package arrow

import "testing"

func TestMinArrows(t *testing.T) {
	st := []struct {
		name   string
		points [][]int
		exp    int
	}{
		{"testcase1", [][]int{[]int{9, 12}, []int{1, 10}, []int{4, 11}, []int{8, 12}, []int{3, 9}, []int{6, 9}, []int{6, 7}}, 2},
		{"testcase2", [][]int{[]int{10, 16}, []int{2, 8}, []int{1, 6}, []int{7, 12}}, 2},
		{"testcase3", [][]int{[]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{7, 8}}, 4},
		{"testcase4", [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{4, 5}}, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minArrows(tt.points)
			if out != tt.exp {
				t.Fatalf("with input points:%v wanted %d but got %d", tt.points, tt.exp, out)
			}
			t.Log("Pass")
		})
	}
}
