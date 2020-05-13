package kd

import "github.com/catorpilor/LeetCode/utils"

func removeKDigits(num string, k int) string {
	return useTraverse(num, k, 0)
}

func useTraverse(num string, k, peak int) string {
	n := len(num)
	if k == 0 {
		idx := 0
		for ; idx < len(num); idx++ {
			if num[idx] == '0' {
				continue
			}
			break
		}
		if idx == len(num) {
			return "0"
		}
		return num[idx:]
	}
	if n == k {
		return "0"
	}

	for ; peak < n-1; peak++ {
		if num[peak] <= num[peak+1] {
			continue
		}
		break
	}
	return useTraverse(num[:peak]+num[peak+1:], k-1, utils.Max(peak-1, 0))
}
