package single

func singleNumber(nums []int) int {
	return useExtra(nums)
}

func useExtra(nums []int) int {
	n := len(nums)
	hset := make(map[int]int, n)
	for _, num := range nums {
		hset[num]++
	}
	for k, v := range hset {
		if v == 1 {
			return k
		}
	}
	return -1
}
