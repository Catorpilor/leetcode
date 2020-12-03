package next

import "testing"

func TestConnect(t *testing.T) {
	st := []struct {
		name      string
		root, exp *TreeNode
	}{
		{"nil node", nil, nil},
		{"testcase1", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}, nil},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := connect(tt.root)
			t.Log("pass")
		})
	}
}
