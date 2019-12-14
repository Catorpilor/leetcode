package cb

import (
	"math"

	"github.com/catorpilor/LeetCode/utils"
)

func assignBikes(workers, bikes [][]int) int {
	nw, nb := len(workers), len(bikes)
	if nw == 0 || nb == 0 {
		return 0
	}
	res := math.MaxInt32
	visited := make([]bool, len(workers))
	dfs(&visited, bikes, workers, 0, 0, &res)
	return res
}

func dfs(visited *[]bool, bikes, workers [][]int, wid, distance int, res *int) {
	if wid >= len(workers) {
		*res = utils.Min(*res, distance)
		return
	}
	if distance > *res {
		return
	}
	for j := 0; j < len(bikes); j++ {
		if (*visited)[j] {
			continue
		}
		(*visited)[j] = true
		dfs(visited, bikes, workers, wid+1, distance+dis(bikes[j], workers[wid]), res)
		(*visited)[j] = false
	}
}

func dis(p1, p2 []int) int {
	return utils.Abs(p1[0]-p2[0]) + utils.Abs(p1[1]-p2[1])
}
