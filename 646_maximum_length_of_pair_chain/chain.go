package chain

import "sort"

func FindLongestChain(pairs [][]int) int {
	l := len(pairs)
	if l == 0 {
		return 0
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	// curMax, maxSofar := 1, 1
	// head, tail := 0, 0 // index of the smallest pair
	// for i := 1; i < len(pairs); i++ {
	// 	if pairs[i][0] > pairs[tail][1] {
	// 		curMax++
	// 		tail = i

	// 	} else {
	// 		if head == tail && pairs[i][1] < pairs[head][1] {
	// 			head = i
	// 			tail = i
	// 			curMax = 1
	// 		}
	// 	}
	// 	if curMax > maxSofar {
	// 		maxSofar = curMax
	// 	}
	// }
	// return maxSofar
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = 1
	}

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	ret := 0
	for i := 0; i < l; i++ {
		if ret < dp[i] {
			ret = dp[i]
		}
	}
	return ret

	// ret := 0
	// for i := 0; i < l; i++ {
	// 	ret++
	// 	curEnd := pairs[i][1]
	// 	for i+1 < l && pairs[i+1][0] <= curEnd {
	// 		i++
	// 	}
	// }
	// return ret
}
