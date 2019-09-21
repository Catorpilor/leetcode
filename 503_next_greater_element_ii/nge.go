package nge

import "github.com/catorpilor/LeetCode/utils"

func nextGreaterElement(nums []int) []int {
	n := len(nums)
	res := make([]int, 2*n)
	for i := range res {
		res[i] = -1
	}
	if n < 1 {
		return res
	}
	expened := make([]int, 0, 2*n)
	expened = append(expened, nums...)
	expened = append(expened, nums[:n-1]...)
	st := utils.NewStack()
	for i := 0; i < 2*n-1; i++ {
		for !st.IsEmpty() && expened[st.Top().(int)] < expened[i] {
			res[st.Pop().(int)] = expened[i]
		}
		st.Push(i)
	}
	return res[:n]
}
