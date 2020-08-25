package ticket

import "testing"

func TestMinCostTickets(t *testing.T) {
	st := []struct {
		name  string
		days  []int
		costs []int
		exp   int
	}{
		{"testcase1", []int{1}, []int{2, 7, 15}, 2},
		{"testcase2", []int{1, 2, 3, 4, 5, 6}, []int{2, 7, 15}, 7},
		{"testcase3", []int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11},
		{"testcase4", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}, 17},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minCostTickets(tt.days, tt.costs)
			if out != tt.exp {
				t.Fatalf("with input days:%v and costs:%v wanted %d but got %d", tt.days, tt.costs, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
