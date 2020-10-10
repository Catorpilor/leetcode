package comp

func bitwiseComp(n int) int {
	return useMask(n)
}

// useMask time complexity O(lgN), space complexity O(1)
func useMask(n int) int {
	mask := 1
	// for example n = 5 -> 101, mask = 111, n's complement is 010
	for n > mask {
		mask = mask<<1 + 1
	}
	return mask - n
}
