package reduce

func numOfSteps(num int) int {
	return useBitOps(num)
}

// useBitOps time complexity O(32), space complexity O(1)
func useBitOps(num int) int {
	var ans int
	for num > 0 {
		if num&1 != 0 {
			num ^= 1
		} else {
			num >>= 1
		}
		ans++
	}
	return ans
}
