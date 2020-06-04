package twocity

import "testing"

func TestMinCost(t *testing.T) {
	st := []struct {
		name  string
		costs [][]int
		exp   int
	}{
		{"testcase1", [][]int{[]int{20, 30}, []int{100, 10}}, 30},
		{"testcase2", [][]int{[]int{10, 20}, []int{30, 200}, []int{400, 50}, []int{30, 20}}, 110},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minCosts(tt.costs)
			if out != tt.exp {
				t.Fatalf("with input costs:%v wanted %d but got %d", tt.costs, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
