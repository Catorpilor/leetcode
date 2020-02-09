package nge

import "github.com/catorpilor/leetcode/utils"

func nextGreaterElement(nums1, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	res := make([]int, n1)
	for i := range res {
		res[i] = -1
	}
	cache := make(map[int]int, n2)
	// all element in nums1, nums2 are unique
	st := utils.NewStack()
	for j := 0; j < n2; j++ {
		for !st.IsEmpty() && st.Top().(int) < nums2[j] {
			cache[st.Pop().(int)] = nums2[j]
		}
		st.Push(nums2[j])
	}
	for i := 0; i < n1; i++ {
		if v, exists := cache[nums1[i]]; exists {
			res[i] = v
		}
	}
	return res
}
