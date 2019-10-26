package nqueue

import (
	"reflect"
	"testing"
)

func TestSolveNQueues(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  [][]string
	}{
		{"n=1", 1, [][]string{[]string{"Q"}}},
		{"n=2", 2, [][]string{}},
		{"n=4", 4, [][]string{[]string{".Q..", "...Q", "Q...", "..Q."}, []string{"..Q.", "Q...", "...Q", ".Q.."}}},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := solveNQueues(tt.n)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input n %d wanted %v but got %v", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
