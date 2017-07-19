package mindepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if l == 0 || r == 0 {
		return l + r + 1
	}
	if l < r {
		return l + 1
	} else {
		return r + 1
	}
}

// my version may be not straight forward as above
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := minDepth2(root.Left)
	r := minDepth2(root.Right)
	ret := l
	if l == 0 {
		ret = r
	}
	if ret > r && r >= 1 {
		ret = r
	}
	return ret + 1
}
