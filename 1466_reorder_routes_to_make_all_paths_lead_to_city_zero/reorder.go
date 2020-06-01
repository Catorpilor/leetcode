package reorder

import "github.com/catorpilor/leetcode/utils"

func minReorder(n int, conns [][]int) int {
	return useDFS(n, conns)
}

// useDFS time complexity O(N), space complexity O(N)
func useDFS(n int, conns [][]int) int {
	// build a tree with adjancy list
	al := make([][]int, n)
	for _, c := range conns {
		from, to := c[0], c[1]
		al[from] = append(al[from], to)
		al[to] = append(al[to], -from)
	}
	visited := make([]bool, n)
	return helper(al, visited, 0)
}

func helper(al [][]int, visited []bool, from int) int {
	var change int
	visited[from] = true
	for _, end := range al[from] {
		if !visited[utils.Abs(end)] {
			alt := 0
			if end > 0 {
				alt++
			}
			change += helper(al, visited, utils.Abs(end)) + alt
		}
	}
	return change
}
