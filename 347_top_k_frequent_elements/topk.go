package topk

// Time Complexity O(n), Space Complexity O(N)
func TopKFreq(nums []int, k int) []int {
	lc := make(map[int]int, len(nums))
	for i := range nums {
		lc[nums[i]]++
	}
	buckets := make(map[int][]int, len(lc))
	maxV := 0
	for k, v := range lc {
		buckets[v] = append(buckets[v], k)
		maxV = max(maxV, v)
	}
	remaining := k
	res := make([]int, 0, k)
	for i := maxV; i > 0; i-- {
		if buckets[i] != nil {
			if len(buckets[i]) >= remaining {
				res = append(res, buckets[i][:remaining]...)
				return res
			}
			remaining -= len(buckets[i])
			res = append(res, buckets[i]...)
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
