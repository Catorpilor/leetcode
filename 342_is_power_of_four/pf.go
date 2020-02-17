package pf

func isPowerOfFour(num int) bool {
	// return useLoops(num)
	// return preCompute(num)
	return useBit(num)
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

// useBit time complexity O(1) space complexity O(1)
func useBit(num int) bool {
	// is power of two is (num > 0) && (num & (num-1) == 0 )
	// we have to eliminate that is power of two but not power of four
	// num = 1 	0000,0001  1-bit at 0
	// num = 4  0000,0100  1-bit at 2
	// num = 16 0001,0000  1-bit at 4
	// num = 64 0100,0000  1-bit at 6
	// num = 2  0000,0010  1-bit at 1
	// num = 8  0000,1000  1-bit at 3
	// num = 32 0010,0000  1-bit at 5
	// num =128 1000,0000  1-bit at 7
	// so that (num & (0xaaaaaaaa) == 0 )
	return (num > 0) && ((num & (num - 1)) == 0) && ((num & 0xaaaaaaaa) == 0)
}
