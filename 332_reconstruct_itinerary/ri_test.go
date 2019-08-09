package ri

import (
	"reflect"
	"testing"
)

func TestConstructIt(t *testing.T) {
	st := []struct {
		name  string
		input [][]string
		exp   []string
	}{
		{"single input", [][]string{[]string{"JFK", "SFO"}}, []string{"JFK", "SFO"}},
		{"leetcode testcase 1", [][]string{[]string{"MUC", "LHR"}, []string{"JFK", "MUC"},
			[]string{"SFO", "SJC"}, []string{"LHR", "SFO"}}, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
		{"testcase 2", [][]string{[]string{"JFK", "A"}, []string{"JFK", "D"}, []string{"A", "C"},
			[]string{"C", "JFK"}, []string{"C", "D"}, []string{"B", "C"}, []string{"D", "B"},
			[]string{"D", "A"}}, []string{"JFK", "A", "C", "D", "B", "C", "JFK", "D", "A"}},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := ConstructIt(tt.input)
			if !reflect.DeepEqual(out, tt.exp) {
				t.Fatalf("with input %v, wanted %v but got %v", tt.input, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
