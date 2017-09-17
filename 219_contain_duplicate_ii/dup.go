package dup

func Dup(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	hm := make(map[int]int)
	for i, c := range nums {
		if v, ok := hm[c]; ok {
			// if exists check first
			if v-1-i >= -k && v-i-1 <= k {
				return true
			}
		}
		// then upate the index
		hm[c] = i + 1
	}
	return false
}
