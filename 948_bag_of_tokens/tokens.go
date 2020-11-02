package tokens

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func maxScore(tokens []int, p int) int {
	return useTwoPointers(tokens, p)
}

// useTwoPointers time complexity O(N*logN), space complexity O(1)
func useTwoPointers(tokens []int, p int) int {
	n := len(tokens)
	sort.Ints(tokens)
	l, r := 0, n-1
	// this is a greedy approach
	var ans, score int
	for l <= r {
		if p >= tokens[l] {
			score++
			p -= tokens[l]
			l++
			ans = utils.Max(ans, score)
		} else if p < tokens[l] && score > 0 {
			p += tokens[r]
			r--
			score--
		} else {
			break
		}
	}
	return ans
}
