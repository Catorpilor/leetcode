package ngl

import (
	"github.com/catorpilor/leetcode/utils"
)

func nextLargeNode(head *utils.ListNode) []int {
	return nlg(head)
}

func sdp(head *utils.ListNode) []int {
	if head == nil {
		return []int{}
	}
	st := utils.NewStack()
	h := head
	for h != nil {
		st.Push(h.Val)
		h = h.Next
	}
	n := st.Len()
	reversd := make([]int, 0, n)
	for !st.IsEmpty() {
		reversd = append(reversd, st.Pop().(int))
	}
	res := make([]int, n)
	dp := make([]int, n)
	dp[0] = -1
	for i := 1; i < n; i++ {
		if reversd[i] < reversd[i-1] {
			dp[i] = i - 1
		} else if reversd[i] == reversd[i-1] {
			dp[i] = dp[i-1]
		} else {
			j := dp[i-1]
			for j != -1 {
				if reversd[j] > reversd[i] {
					dp[i] = j
					break
				}
				j = dp[j]
			}
			if j == -1 && dp[i] == 0 {
				dp[i] = j
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		if dp[i] != -1 {
			res[n-1-i] = reversd[dp[i]]
		}
	}

	return res
}

func nlg(head *utils.ListNode) []int {
	if head == nil {
		return []int{}
	}
	var arr []int
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p.Val)
	}
	n := len(arr)
	res := make([]int, n)
	st := utils.NewStack()
	for i := 0; i < n; i++ {
		for !st.IsEmpty() && arr[st.Top().(int)] < arr[i] {
			idx := st.Pop().(int)
			res[idx] = arr[i]
		}
		st.Push(i)
	}
	return res
}
