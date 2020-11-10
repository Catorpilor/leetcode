package seat

import "github.com/catorpilor/leetcode/utils"

func maxDist(seats []int) int {
	return useStraight(seats)
}

// useStraight time complexity O(N), space complexity O(1)
func useStraight(seats []int) int {
	n := len(seats)
	var ans, leftZero, prevOne int
	leftZero, prevOne = -1, -1
	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			if leftZero != -1 {
				var cur int
				if prevOne == -1 {
					cur = i - leftZero
				} else {
					cur = (i - prevOne) >> 1
				}
				ans = utils.Max(ans, cur)
			}
			leftZero = -1
			prevOne = i
			continue
		}
		// seats[i] = 0 here
		if leftZero == -1 {
			leftZero = i
		}
	}
	if leftZero != -1 {
		// seats end with 0, for example [1,0,0,0,0]
		ans = utils.Max(ans, n-1-prevOne)
	}
	return ans
}
