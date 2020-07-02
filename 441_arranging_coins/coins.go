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

// useBinarySearch time complexity O(lgN), space complexity O(1)
func useBinarySearch(n int) int {
	l, r := 0, n
	var ans int
	for l <= r {
		ans = l + (r-l)/2
		total := ans * (ans + 1) / 2
		if total == n {
			return ans
		} else if total < n {
			l = ans + 1
		} else {
			r = ans - 1
		}
	}
	return r
}
