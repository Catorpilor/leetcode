package nge

import "github.com/catorpilor/leetcode/utils"

func nextGreaterElement(nums []int) []int {
	// return expended(nums)
	return circular(nums)
}

func expended(nums []int) []int {
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

func circular(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	st := utils.NewStack()
	for i := 0; i < 2*n; i++ {
		num := nums[i%n]
		for !st.IsEmpty() && nums[st.Top().(int)] < num {
			res[st.Pop().(int)] = num
		}
		if i < n {
			st.Push(i)
		}
	}
	return res
}
