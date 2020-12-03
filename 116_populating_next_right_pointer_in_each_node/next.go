package next

type TreeNode struct {
	Val               int
	Left, Right, Next *TreeNode
}

func connect(root *TreeNode) *TreeNode {
	return uselevelorder(root)
}

// uselevelorder time complexity O(N), space complexity O(N) which is not constant.
func uselevelorder(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	}
	var st []*TreeNode
	st = append(st, root)
	nst := len(st)
	for nst != 0 {
		var tmp []*TreeNode
		for i, node := range st {
			if node.Left != nil && node.Right != nil {
				tmp = append(tmp, node.Left, node.Right)
			}
			if i+1 < nst {
				node.Next = st[i+1]
			}
		}
		st = tmp
		nst = len(st)
	}
	return root
}
