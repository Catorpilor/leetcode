package pf

func isPowerOfFour(num int) bool {
	return useLoops(num)
}

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
