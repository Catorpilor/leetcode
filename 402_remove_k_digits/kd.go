package kd

import (
	"github.com/catorpilor/leetCode/utils"
)

func removeKDigits(num string, k int) string {
	return useTraverse(num, k, 0)
}

// useTraverse time complexity O(KN), space complexity O(KN)
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

// useStack time complexity O(N), space complexity O(N)
func useStack(num string, k int) string {
	n := len(num)
	st := make([]byte, 0, n)
	left := n - k
	for _, c := range num {
		nst := len(st)
		for nst > 0 && st[nst-1] > byte(c) && k > 0 {
			st = st[:nst-1]
			nst--
			k--
		}
		st = append(st, byte(c))
	}
	var idx int
	// fmt.Println(string(st))
	for ; idx < left; idx++ {
		if st[idx] == '0' {
			continue
		}
		break
	}
	if idx == left {
		return "0"
	}
	// fmt.Println(idx, left)
	return string(st[idx:left])
}
