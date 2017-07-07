package traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrderBottom traverse the binary tree by level and print the nodes in the reverse order
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var curLevel []int
		var levelNodes []*TreeNode
		for _, n := range queue {
			curLevel = append(curLevel, n.Val)
			if n.Left != nil {
				levelNodes = append(levelNodes, n.Left)
			}
			if n.Right != nil {
				levelNodes = append(levelNodes, n.Right)
			}
		}
		queue = levelNodes
		ret = append(ret, curLevel)
	}

	// reverse
	for i, j := 0, len(queue)-1; i < j; i, j = i+1, j-1 {
		queue[i], queue[j] = queue[j], queue[i]
	}
}
