package pb

import "testing"

func TestPossibleBipartition(t *testing.T) {
	st := []struct {
		name string
		n    int
		dis  [][]int
		exp  bool
	}{
		{"n=1", 1, nil, true},
		{"testcase1", 4, [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 4}}, true},
		{"testcase2", 3, [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 3}}, false},
		{"testcase5", 5, [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{4, 5}, []int{1, 5}}, false},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := possibleBipartition(tt.n, tt.dis)
			if out != tt.exp {
				t.Fatalf("with input N:%d and dis:%v wanted %t but got %t", tt.n, tt.dis, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
