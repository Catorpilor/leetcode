package single

func singleNumber(nums []int) int {
	// return useExtra(nums)
	return useBitOp(nums)
}

// useExtra time complexity O(N), space complexity O(N)
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

// useBitOp time complexity O(N), space complexity O(1)
func useBitOp(nums []int) int {
	// k = 3 => 11 => k1 = 1, k2 = 1
	// m = logk = 2 use two variables
	// m > logk so we use a mask
	// mask = ^(y1 & y2) where y1=x1 if k1==1, y2=~x2 if k2 == 0
	var x1, x2, mask int
	for _, num := range nums {
		x2 ^= (x1 & num)
		x1 ^= num
		mask = ^(x1 & x2)
		x1 &= mask
		x2 &= mask
	}
	return x1
}
