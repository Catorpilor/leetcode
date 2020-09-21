package image

import (
	"fmt"

	"github.com/catorpilor/leetcode/utils"
)

func largestOverlap(a, b [][]int) int {
	return useHashmap(a, b)
}

// useHashmap time complexity O(N^4) worst cases, space compelxity O(N^2)
func useHashmap(a, b [][]int) int {
	m := len(a)
	if m < 1 {
		return 0
	}
	n := len(a[0])
	if n < 1 {
		return 0
	}
	// create two lists save the pixel coordinates
	la, lb := make([][]int, 0, m*n), make([][]int, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 1 {
				la = append(la, []int{i, j})
			}
			if b[i][j] == 1 {
				lb = append(lb, []int{i, j})
			}
		}
	}
	// create a map to map vector (from pixel in A to a pixel in B) to its count
	vs := make(map[string]int, m*n)
	for _, pa := range la {
		for _, pb := range lb {
			key := fmt.Sprintf("%d %d", pa[0]-pb[0], pa[1]-pb[1])
			vs[key]++
		}
	}
	var ans int
	for _, v := range vs {
		ans = utils.Max(ans, v)
	}
	return ans
}
