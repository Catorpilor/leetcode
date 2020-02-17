package pf

func isPowerOfFour(num int) bool {
	// return useLoops(num)
	return preCompute(num)
}

// useLoops time complexity O(lgN), space complexity O(1)
func useLoops(num int) bool {
	if num <= 0 {
		return false
	}
	for num != 1 {
		if num%4 != 0 {
			return false
		}
		num >>= 2
	}
	return true
}

var results map[int]bool

func init() {
	if results == nil {
		results = make(map[int]bool)
	}
	// since num <= 2^31 - 1, so the max power of 4 is 4^15
	results[1] = true
	base := 1
	for i := 1; i <= 15; i++ {
		base *= 4
		results[base] = true
	}
}

// preCompute time complexity O(1), space complexity O(1)
func preCompute(num int) bool {
	return results[num]
}
