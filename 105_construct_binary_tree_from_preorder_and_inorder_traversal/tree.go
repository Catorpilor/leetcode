package tree

import "github.com/catorpilor/leetcode/utils"

func buildTree(preorder, inorder []int) *utils.TreeNode {
	return useDFS(preorder, inorder)
}

// useDFS time complexity O(N) space complexity  O(1)
func useDFS(preorder, inorder []int) *utils.TreeNode {
	np, ni := len(preorder), len(inorder)
	if np < 1 && ni < 1 {
		return nil
	}
	val := preorder[0] // root val
	root := &utils.TreeNode{Val: val}
	if np == 1 && ni == 1 {
		return root
	}
	var pos int // find root's pos in inorder array
	for pos < ni {
		if inorder[pos] == val {
			break
		}
		pos++
	}
	root.Left = useDFS(preorder[1:1+pos], inorder[:pos])
	root.Right = useDFS(preorder[1+pos:], inorder[pos+1:])
	return root
}
