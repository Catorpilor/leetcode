package tree

import (
	"github.com/catorpilor/leetcode/utils"
)

func buildTree(inorder, postorder []int) *utils.TreeNode {
	return useDFS(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

// useDFS time complexity O(N), space complexity O(1) ?
func useDFS(inorder, postorder []int, li, ri, lp, rp int) *utils.TreeNode {
	if li > ri || lp > rp {
		return nil
	}

	root := &utils.TreeNode{Val: postorder[rp]}
	pos := li
	for pos <= ri {
		if inorder[pos] == root.Val {
			break
		}
		pos++
	}
	leftNodes := pos - li
	// fmt.Printf("li:%d, ri:%d, lp:%d, rp:%d, pos:%d, lefeNode:%d\n", li, ri, lp, rp, pos, leftNodes)
	root.Left = useDFS(inorder, postorder, li, pos-1, lp, lp+leftNodes-1)
	root.Right = useDFS(inorder, postorder, pos+1, ri, lp+leftNodes, rp-1)
	return root
}
