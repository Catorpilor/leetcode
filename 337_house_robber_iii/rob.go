package rob

import (
	"github.com/catorpilor/leetcode/utils"
)

func Rob(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	// 1. root is robbed
	// 2. root is not robbed.
	var ret int
	if root.Left != nil {
		ret += Rob(root.Left.Left) + Rob(root.Left.Right)
	}
	if root.Right != nil {
		ret += Rob(root.Right.Left) + Rob(root.Right.Right)
	}
	return utils.Max(root.Val+ret, Rob(root.Left)+Rob(root.Right))
}

func Rob2(root *utils.TreeNode) int {
	mmap := make(map[*utils.TreeNode]int)
	return robSub(root, mmap)
}

func robSub(node *utils.TreeNode, hm map[*utils.TreeNode]int) int {
	if node == nil {
		return 0
	}
	if v, ok := hm[node]; ok {
		return v
	}
	var ret int
	if node.Left != nil {
		ret += robSub(node.Left.Left, hm) + robSub(node.Left.Right, hm)
	}
	if node.Right != nil {
		ret += robSub(node.Right.Left, hm) + robSub(node.Right.Right, hm)
	}
	ret = utils.Max(ret+node.Val, robSub(node.Left, hm)+robSub(node.Right, hm))
	hm[node] = ret
	return ret
}

func Rob3(root *utils.TreeNode) int {
	res := robSub2(root)
	// res[0] represent root is not robbed
	// res[1] represent root is robbed
	return utils.Max(res[0], res[1])
}

func robSub2(node *utils.TreeNode) []int {
	res := []int{0, 0}
	if node == nil {
		return res
	}
	left := robSub2(node.Left)
	right := robSub2(node.Right)

	res[0] = utils.Max(left[0], left[1]) + utils.Max(right[0], right[1])
	res[1] = node.Val + left[0] + right[0]
	return res
}

// type NodeWithLevel struct {
// 	node  *utils.TreeNode
// 	level int
// }

// func helper(node *utils.TreeNode, ans *int) {
// 	// level order travesal
// 	if node == nil {
// 		return
// 	}
// 	q := utils.NewQueue()
// 	q.Enroll(NodeWithLevel{node: node})
// 	var res [2048]int
// 	for !q.IsEmpty() {
// 		n := q.Pull().(NodeWithLevel)
// 		// fmt.Println(n.level, n.node.Val)
// 		if n.node.Left != nil {
// 			q.Enroll(NodeWithLevel{
// 				node:  n.node.Left,
// 				level: n.level + 1,
// 			})
// 		}
// 		if n.node.Right != nil {
// 			q.Enroll(NodeWithLevel{
// 				node:  n.node.Right,
// 				level: n.level + 1,
// 			})
// 		}
// 		res[n.level] += n.node.Val
// 	}
// 	var even, odd int
// 	for i := 0; i < 20; i++ {
// 		if res[i] == 0 {
// 			break
// 		}
// 		if i%2 == 0 {
// 			even += res[i]
// 		} else {
// 			odd += res[i]
// 		}
// 	}
// 	*ans = utils.Max(even, odd)
// }
