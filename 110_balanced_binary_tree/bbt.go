package bbt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// topdown approach

func isBalancedTD(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := depth(root.Left)
	r := depth(root.Right)

	return (l-r <= 1 && l-r >= -1) && isBalancedTD(root.Left) && isBalancedTD(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left)
	r := depth(root.Right)
	ret := l
	if r > ret {
		ret = r
	}
	return ret + 1
}

// bottom up
func isBalancedBU(root *TreeNode) bool {
	return dfsHeight(root) != -1
}

func dfsHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfsHeight(root.Left)
	if l == -1 {
		return -1
	}
	r := dfsHeight(root.Right)
	if r == -1 {
		return -1
	}
	if (l-r > 1) || (l-r < -1) {
		return -1
	}
	ret := l
	if r > ret {
		ret = r
	}
	return ret + 1
}
