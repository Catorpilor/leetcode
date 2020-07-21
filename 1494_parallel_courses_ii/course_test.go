package course

import "testing"

func TestMinimumSems(t *testing.T) {
	st := []struct {
		name string
		n    int
		deps [][]int
		k    int
		exp  int
	}{
		{"no deps", 5, nil, 2, 3},
		{"testcase1", 4, [][]int{[]int{2, 1}, []int{3, 1}, []int{1, 4}}, 2, 3},
		{"testcase2", 5, [][]int{[]int{2, 1}, []int{3, 1}, []int{4, 1}, []int{1, 5}}, 2, 4},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minSems(tt.n, tt.deps, tt.k)
			if out != tt.exp {
				t.Fatalf("with input n:%d, deps:%v and k:%d wanted %d but got %d", tt.n, tt.deps, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
