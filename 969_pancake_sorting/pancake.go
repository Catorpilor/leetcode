package pancake

func pancakeSorting(arr []int) []int {
	return useStraight(arr)
}

// useStraight time complexity O(N^2), space complexity O(1)
func useStraight(arr []int) []int {
	n := len(arr)
	// this uses the condition of the permutation of [1..n]
	// if not we could sort the arr first and traverse from the tail.
	ans := make([]int, 0, 10*n)
	for x := n; x > 0; x-- {
		var i int
		for arr[i] != x {
			i++
		}
		reverse(arr, i+1)
		ans = append(ans, i+1)
		reverse(arr, x)
		ans = append(ans, x)
	}
	return ans
}

func reverse(arr []int, k int) {
	l, r := 0, k-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
