package sum

import "github.com/catorpilor/leetcode/utils"

func pathSum(root *utils.TreeNode) int {
	return useBackTracking(root)
}

// useBackTracking time compelxity O(N), space complexity O(log(N))
func useBackTracking(root *utils.TreeNode) int {
	var ans int
	helper(root, &ans, []int{})
	return ans
}

func helper(root *utils.TreeNode, ans *int, path []int) {
	if root == nil {
		return
	}
	np := len(path)
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		// it's a leaf node
		*ans += genNum(path)
		path = path[:np]
		return
	}
	helper(root.Left, ans, path)
	helper(root.Right, ans, path)
}

func genNum(nums []int) int {
	var ans int
	n := len(nums)
	for i := 0; i < n; i++ {
		ans += nums[i] * (1 << (n - 1 - i))
	}
	return ans
}

// useDfs time complexity O(N) space complexity O(log(N))
func useDfs(root *utils.TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *utils.TreeNode, val int) int {
	if root == nil {
		return 0
	}
	val = val*2 + root.Val
	if root.Left == nil && root.Right == nil {
		return val
	}
	return dfs(root.Left, val) + dfs(root.Right, val)
}
