package course

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCheckIfPrerequisite(t *testing.T) {
	st := []struct {
		name    string
		n       int
		pres    [][]int
		queries [][]int
		exp     []bool
	}{
		{"testcase1", 2, [][]int{[]int{1, 0}}, [][]int{[]int{0, 1}, []int{1, 0}}, []bool{false, true}},
		{"testcase2", 2, [][]int{}, [][]int{[]int{0, 1}, []int{1, 0}}, []bool{false, false}},
		{"testcase3", 3, [][]int{[]int{1, 2}, []int{1, 0}, []int{2, 0}}, [][]int{[]int{1, 0}, []int{1, 2}}, []bool{true, true}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := checkIfPrerequisite(tt.n, tt.pres, tt.queries)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("with input n:%d, pres:%v and queries:%v  mismatch (-want +got):\n%s", tt.n, tt.pres, tt.queries, diff)
			}
			t.Log("pass")
		})
	}
}
