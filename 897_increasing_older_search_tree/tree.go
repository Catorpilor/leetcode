package tree

import "github.com/catorpilor/leetcode/utils"

func increasingBST(root *utils.TreeNode) *utils.TreeNode {
	// return useDFS(root)
	return useDFS2(root, nil)
}

// useDFS time complexity O(N), space complexity O(lgN)
func useDFS(root *utils.TreeNode) *utils.TreeNode {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	}
	var lh, lt *utils.TreeNode
	if root.Left != nil {
		lh, lt = helper(root.Left)
	}
	var rh *utils.TreeNode
	if root.Right != nil {
		rh, _ = helper(root.Right)
	}
	if lt != nil {
		lt.Right = root
	} else {
		// root.Left is nil, so root remains the root
		lh = root
	}
	// rh, _ := helper(root.Right)
	root.Left = nil
	root.Right = rh
	return lh
}

func helper(root *utils.TreeNode) (head, tail *utils.TreeNode) {
	if root.Left == nil {
		head = root
	} else {
		var lt *utils.TreeNode
		head, lt = helper(root.Left)
		lt.Right = root
		// rearrange the tree to avoid TLE
		root.Left = nil
	}
	if root.Right != nil {
		lh, lt := helper(root.Right)
		root.Right = lh
		tail = lt
	} else {
		tail = root
	}
	return
}

// useDFS2 time complexity O(N), space complexity O(logN).
func useDFS2(root, tail *utils.TreeNode) *utils.TreeNode {
	// inorder traversal.
	if root == nil {
		return tail
	}
	res := useDFS2(root.Left, root)
	root.Left = nil
	root.Right = useDFS2(root.Right, tail)
	return res
}
