package ps

import "github.com/catorpilor/leetcode/utils"

func NumSquares(num int) int {
	if num <= 0 {
		return 0
	}
	ns := make([]int, num+1)
	for i := range ns {
		ns[i] = i
	}
	i, j := 2, 0
	for i*i <= num && j <= num {
		if i*i <= j {
			ns[j] = utils.Min(ns[j], ns[j-i*i]+1)
		}
		if j == num {
			j, i = 0, i+1
		} else {
			j += 1
		}
	}
	return ns[num]
}
