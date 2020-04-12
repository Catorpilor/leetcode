package candy

import (
	"github.com/catorpilor/leetcode/utils"
)

func numOfCandies(ratings []int) int {
	// return useBruteForce(ratings)
	return useTwoArray(ratings)
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

// useTwoArray time complexity O(N), space complexity O(N)
func useTwoArray(ratings []int) int {
	n := len(ratings)
	if n <= 1 {
		return n
	}
	leftTurn, rightTurn := make([]int, n), make([]int, n)
	leftTurn[0], rightTurn[n-1] = 1, 1
	for i := 1; i < n; i++ {
		if ratings[i] <= ratings[i-1] {
			leftTurn[i] = 1
		} else {
			leftTurn[i] = leftTurn[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] <= ratings[i+1] {
			rightTurn[i] = 1
		} else {
			rightTurn[i] = rightTurn[i+1] + 1
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		ans += utils.Max(leftTurn[i], rightTurn[i])
	}
	return ans
}
