package dtime

import "fmt"

func maxTime(arr []int) string {
	return usePermutation(arr)
}

// usePermutation time complexity O(1), space complexity O(1)
func usePermutation(arr []int) string {
	var hour, mint int
	curMax := -1
	for _, p := range genPermutations(arr) {
		h := p[0]*10 + p[1]
		if h > 23 {
			continue
		}
		mm := p[2]*10 + p[3]
		if mm > 59 {
			continue
		}
		lc := h*100 + mm
		if lc > curMax {
			curMax = lc
			hour = h
			mint = mm
		}
	}
	if curMax == -1 {
		return ""
	}
	return fmt.Sprintf("%02d:%02d", hour, mint)
}

func genPermutations(xs []int) [][]int {
	permuts := make([][]int, 0, 24)
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == 4 {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}
