package list

import "github.com/catorpilor/leetcode/utils"

func numComps(head *utils.ListNode, g []int) int {
	return iterator(head, g)
}

// iterator time complexity O(N+len(g)), space compleixty O(len(g))
func iterator(head *utils.ListNode, g []int) int {
	gset := make(map[int]bool, len(g))
	for i := range g {
		gset[g[i]] = true
	}
	var ret int
	for head != nil {
		if gset[head.Val] && (head.Next == nil || !gset[head.Next.Val]) {
			ret++
		}
		head = head.Next
	}
	return ret
}
