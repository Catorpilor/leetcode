package div

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCalcEquation(t *testing.T) {
	st := []struct {
		name      string
		equations [][]string
		values    []float64
		queries   [][]string
		exp       []float64
	}{
		{"testcase1", [][]string{[]string{"a", "b"}, []string{"b", "c"}}, []float64{2.0, 3.0}, [][]string{[]string{"a", "c"}, []string{"b", "a"}, []string{"a", "e"}, []string{"a", "a"}, []string{"x", "x"}},
			[]float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}},
		{"testcase2", [][]string{[]string{"a", "b"}, []string{"b", "c"}, []string{"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{[]string{"a", "c"}, []string{"c", "b"}, []string{"bc", "cd"}, []string{"cd", "bc"}},
			[]float64{3.75000, 0.40000, 5.00000, 0.20000}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := calcEquation(tt.equations, tt.values, tt.queries)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
