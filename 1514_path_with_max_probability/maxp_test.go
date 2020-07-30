package maxp

import "testing"

func TestMaxProb(t *testing.T) {
	st := []struct {
		name       string
		n          int
		edges      [][]int
		succProb   []float64
		start, end int
		exp        float64
	}{
		{"testcase1", 3, [][]int{[]int{0, 1}, []int{1, 2}, []int{0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2, 0.25000},
		{"testcase2", 3, [][]int{[]int{0, 1}, []int{1, 2}, []int{0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2, 0.3000},
		{"testcase3", 3, [][]int{[]int{0, 1}}, []float64{0.5}, 0, 2, 0.000},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxProb(tt.n, tt.edges, tt.succProb, tt.start, tt.end)
			if out != tt.exp {
				t.Fatalf("wanted %03f but got %03f", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
