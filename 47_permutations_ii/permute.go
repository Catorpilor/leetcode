package permute

func permuteUnique(nums []int) [][]int {
	var res [][]int
	if nums == nil || len(nums) == 0 {
		return res
	}

	hset := make(map[int]int)
	for i := range nums {
		hset[nums[i]]++
	}
	res = permute(res, hset, []int{}, len(nums))
	return res

}

func permute(res [][]int, hset map[int]int, r []int, n int) [][]int {
	if n <= 0 {
		res = append(res, r)
		return res
	}
	for k := range hset {
		if hset[k] > 0 {
			hset[k]--
			r = append(r, k)
			res = permute(res, hset, r, n-1)
			hset[k]++
		}
	}
	return res
}
