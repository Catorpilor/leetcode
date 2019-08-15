package oab

import "github.com/catorpilor/leetcode/utils"

func MinTransfers(transactions [][]int) int {
	lcs := make(map[int]int)
	for _, tt := range transactions {
		lcs[tt[0]] -= tt[2]
		lcs[tt[1]] += tt[2]
	}
	debts := make([]int, 0, len(lcs))
	for _, v := range lcs {
		debts = append(debts, v)
	}
	return settle(0, debts)
}

func settle(start int, debt []int) int {
	for start < len(debt) && debt[start] == 0 {
		start++
	}
	if start == len(debt) {
		return 0
	}
	r := 1<<31 - 1
	for i := start + 1; i < len(debt); i++ {
		if debt[i]*debt[start] < 0 { // skip same sign debt
			debt[i] += debt[start]
			r = utils.Min(r, 1+settle(start+1, debt))
			debt[i] -= debt[start]
		}
	}
	return r
}
