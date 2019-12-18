package cb

import (
	"container/heap"

	"github.com/catorpilor/LeetCode/utils"
)

func assignBikes(workers, bikes [][]int) []int {
	nw, nb := len(workers), len(bikes)
	if nw == 0 || nb == 0 {
		return []int{}
	}
	tupq := utils.TuplePriorityQueue{}
	// loop through every pairs of bikes and workers
	for i := 0; i < nw; i++ {
		w := workers[i]
		for j := 0; j < nb; j++ {
			b := bikes[j]
			dist := utils.Abs(w[0]-b[0]) + utils.Abs(w[1]-b[1])
			heap.Push(&tupq, [3]int{dist, i, j})
		}
	}
	res := make([]int, nb)
	// init res to -1 labled as none visited
	for i := range res {
		res[i] = -1
	}

	set := make(map[int]bool, nb)
	for len(set) < nb {
		tup := heap.Pop(&tupq).([3]int)
		if res[tup[1]] == -1 && !set[tup[2]] {
			res[tup[1]] = tup[2]
			set[tup[2]] = true
		}
	}
	return res
}
