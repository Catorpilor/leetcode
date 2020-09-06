package all

import "github.com/catorpilor/leetcode/utils"

func allElements(root1, root2 *utils.TreeNode) []int {
	return useMorris(root1, root2)
}

// useMorris time complexity O(N), space complexity O(N)
func useMorris(root1, root2 *utils.TreeNode) []int {
	u1, u2 := utils.InorderTraversal(root1), utils.InorderTraversal(root2)
	n1, n2 := len(u1), len(u2)
	ans := make([]int, 0, n1+n2)
	var i, j int
	for i != n1 && j != n2 {
		if u1[i] <= u2[j] {
			ans = append(ans, u1[i])
			i++
		} else {
			ans = append(ans, u2[j])
			j++
		}
	}
	if i != n1 {
		ans = append(ans, u1[i:]...)
	}
	if j != n2 {
		ans = append(ans, u2[j:]...)
	}
	return ans
}
