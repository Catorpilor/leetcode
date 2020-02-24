package convert

import "github.com/catorpilor/leetcode/utils"

func decimalValue(head *utils.ListNode) int {
	return iterator(head)
}

// iterator time complexity O(N), space complexity O(1)
// follow the basic implementation of golang's strconv.ParseUInt
func iterator(head *utils.ListNode) int {
	var ret int
	cur := head
	for cur != nil {
		// ret *= base here base = 2
		ret <<= 1
		ret += cur.Val
		cur = cur.Next
	}
	return ret
}

func useOr(head *utils.ListNode) int {
	var ret int
	cur := head
	for cur != nil {
		ret = (ret << 1) | cur.Val
		cur = cur.Next
	}
	return ret
}
