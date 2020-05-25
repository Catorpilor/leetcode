package path

import "github.com/catorpilor/leetcode/utils"

func maxNumOfPalindromePath(root *utils.TreeNode) int {
	return useBackTracking(root)
}

// useBackTracking time complexity O(N),space complexity O(logN)
func useBackTracking(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	// pathSet stores the occurrences of each value in the path
	pathSet := make(map[int]int, 9)
	var ans int
	helper(&ans, root, pathSet)
	return ans
}

func helper(ans *int, root *utils.TreeNode, path map[int]int) {
	path[root.Val]++
	if root.Left == nil && root.Right == nil {
		// path end here
		if canBeAPalindrome(path) {
			(*ans)++
		}
	}
	if root.Left != nil {
		helper(ans, root.Left, path)
	}
	if root.Right != nil {
		helper(ans, root.Right, path)
	}
	path[root.Val]--
	if path[root.Val] == 0 {
		delete(path, root.Val)
	}
}

func canBeAPalindrome(path map[int]int) bool {
	var total, oddOnes int
	for _, v := range path {
		if v&1 != 0 {
			oddOnes++
		}
		total += v
	}
	return total&1 != 0 && oddOnes == 1 || total&1 == 0 && oddOnes == 0
}
