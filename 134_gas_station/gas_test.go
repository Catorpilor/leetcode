package gas

import "testing"

func TestCanComplete(t *testing.T) {
	st := []struct {
		name string
		gas  []int
		cost []int
		exp  int
	}{
		{"testcase1", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{"testcase2", []int{2, 3, 4}, []int{3, 4, 3}, -1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := canComplete(tt.gas, tt.cost)
			if out != tt.exp {
				t.Fatalf("with input gas:%v and costs: %v wanted %d but got %d", tt.gas, tt.cost, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
