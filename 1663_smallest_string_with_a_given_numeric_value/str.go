package str

import "github.com/catorpilor/leetcode/utils"

func getSmallestString(n, k int) string {
	return useGreedy(n, k)
}

func useGreedy(n, k int) string {
	sb := make([]byte, n)
	for i := range sb {
		sb[i] = 'a'
	}
	k -= n
	i := n - 1
	for k > 0 {
		tmp := utils.Min(k, 25)
		sb[i] = sb[i] + byte(tmp)
		k -= tmp
		i--
	}
	return string(sb)
}
