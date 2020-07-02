package coins

func arrangeCoins(n int) int {
	return useTraverse(n)
}

// usetraverse time complexity O(N), space complexity O(1)
func useTraverse(n int) int {
	if n < 1 {
		return n
	}
	ans := 1
	for n > 0 {
		n -= ans
		if n >= ans+1 {
			ans++
		}
	}
	return ans
}
