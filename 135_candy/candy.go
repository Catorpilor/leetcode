package candy

func numOfCandies(ratings []int) int {
	return useBruteForce(ratings)
}

// useBruteForce time complexity O(N^2), space complexity O(N)
func useBruteForce(ratings []int) int {
	n := len(ratings)
	if n <= 1 {
		return n
	}
	flag := true
	store := make([]int, n)
	for i := range store {
		store[i] = 1
	}
	for flag {
		flag = false
		for i := 0; i < n; i++ {
			if i > 0 && ratings[i] > ratings[i-1] && store[i] <= store[i-1] {
				store[i] = store[i-1] + 1 // update need to validate
				flag = true
			}
			if i < n-1 && ratings[i] > ratings[i+1] && store[i] <= store[i+1] {
				store[i] = store[i+1] + 1 // update need to validate
				flag = true
			}
		}
	}
	var ans int
	for i := range store {
		ans += store[i]
	}
	return ans
}
