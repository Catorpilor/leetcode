package pow

func powN(x float64, n int) float64 {
	return useRecursion(x, n)
}

func useRecursion(x float64, n int) float64 {
	return helper(x, n)
}

// helper time complexity O(logN) space complexity O(1)
func helper(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	ans := 1.0
	for n != 0 {
		if n&1 != 0 {
			if n < 0 {
				ans /= x
			} else {
				ans *= x
			}
		}
		x *= x
		n /= 2
	}
	return ans
}
