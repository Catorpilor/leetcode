package chips

import "testing"

func TestMinCost(t *testing.T) {
	st := []struct {
		name    string
		postion []int
		exp     int
	}{
		{"testcase1", []int{1, 2, 3}, 1},
		{"testcase2", []int{2, 2, 2, 3, 3}, 2},
		{"testcase3", []int{1, 100000}, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minCost(tt.postion)
			if out != tt.exp {
				t.Fatalf("with input position:%v wanted %d but got %d", tt.postion, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
