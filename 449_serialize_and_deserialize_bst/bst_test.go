package bst

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
	"github.com/google/go-cmp/cmp"
)

func initWithT(t *testing.T) *Codec {
	cdc := Constructor()

	t.Cleanup(func() {
		cdc = nil
	})
	return cdc
}

func TestSandDBST(t *testing.T) {
	cdc := initWithT(t)
	st := []struct {
		name string
		root *utils.TreeNode
		ds   string
	}{
		{"testcase1", utils.ConstructTree([]int{2, 1, 3}), "213"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := cdc.serialize(tt.root)
			if diff := cmp.Diff(tt.ds, out); diff != "" {
				t.Fatalf("(-wanted, +got), %s", diff)
			}
			t.Log("pass")
		})
	}

}
