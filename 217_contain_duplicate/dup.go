package dup

func Dup(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	hm := make(map[int]bool)
	for _, c := range nums {
		if _, ok := hm[c]; ok {
			return true
		} else {
			hm[c] = true
		}
	}
	return false
}
