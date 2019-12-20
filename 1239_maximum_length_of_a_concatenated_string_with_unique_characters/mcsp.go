package mcsp

import "github.com/catorpilor/leetcode/utils"

func maxLength(arr []string) int {
	return bitSets(arr)
}

func bitSets(arr []string) int {
	var res int
	dp := make([]int, 0, len(arr))
	dp = append(dp, 0)
	for _, s := range arr {
		var a, dup int
		for i := range s {
			dup |= a & (1 << (s[i] - 'a'))
			a |= 1 << (s[i] - 'a')
		}
		if dup > 0 {
			continue
		}
		for i := len(dp) - 1; i >= 0; i-- {
			if a&dp[i] > 0 {
				continue
			}
			dp = append(dp, a|dp[i])
			res = utils.Max(res, utils.NumOfOnes(a|dp[i]))
		}
	}
	return res
}

/*
func bruteForce(arr []string) int {
	var res int
	n := len(arr)
	bmap := 1 << 26
	for i := 0; i < n-1; i++ {
		innLen := len(arr[i])
		res = utils.Max(res, innLen)
		bmap = set(bmap, arr[i])
		for j := i + 1; j < n; j++ {
			var bFlag bool
			for k := 0; k < len(arr[j]); k++ {
				pos := arr[j][k] - 'a'
				if bmap&(1<<pos) != 0 {
					bFlag = true
					break
				}
			}
			if !bFlag {
				innLen += len(arr[j])
			}
		}
		res = utils.Max(res, innLen)
		bmap = clr(bmap, arr[i])
	}
	return res
}

func set(bmap int, s string) int {
	for i := range s {
		k := s[i] - 'a'
		bmap |= 1 << k
	}
	return bmap
}

func clr(bmap int, s string) int {
	for i := range s {
		k := s[i] - 'a'
		bmap ^= 1 << k
	}
	return bmap
}
*/
