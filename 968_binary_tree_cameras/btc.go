package btc

import "github.com/catorpilor/leetcode/utils"

func MinCameras(root *utils.TreeNode) int {
	var res int
	// state := dfs(root, &res)
	state := simpleDFS(root, &res)
	if state < 1 {
		// need camera
		res++
	}
	return res
}

// dfs depth-first-search the tree
// use greedy algorithem
// put no cameras on leaves
// return values
// 0 means child need to be covered.
// 1 means node with camera, covered.
// 2 means parent node already covered by node.

func dfs(root *utils.TreeNode, res *int) int {
	var neededCamera, covered int
	if root.Left == nil && root.Right == nil {
		return 0
	}
	if root.Left != nil {
		state := dfs(root.Left, res)
		if state == 0 {
			neededCamera = 1
			covered = 1
		} else if state == 1 {
			covered = 1
		}
	}
	if root.Right != nil {
		state := dfs(root.Right, res)
		if state == 0 {
			neededCamera = 1
			covered = 1
		} else if state == 1 {
			covered = 1
		}
	}
	if neededCamera > 0 {
		*res++
		return 1
	}
	if covered > 0 {
		return 2
	}
	return 0
}

func simpleDFS(root *utils.TreeNode, res *int) int {
	if root == nil {
		return 2
	}
	l, r := simpleDFS(root.Left, res), simpleDFS(root.Right, res)
	if l == 0 || r == 0 {
		*res++
		return 1
	}
	if l == 1 || r == 1 {
		return 2
	}
	return 0
}
