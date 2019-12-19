package cb

import (
	"container/heap"

	"github.com/catorpilor/LeetCode/utils"
)

func assignBikes(workers, bikes [][]int) []int {
	return pq(workers, bikes)
}

func pq(workers, bikes [][]int) []int {
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

// bucket use bucket sort to solve the problem
func bucket(workers, bikes [][]int) []int {
	nw, nb := len(workers), len(bikes)
	if nw == 0 || nb == 0 {
		return []int{}
	}
	// since distance ranges from [0,2000)
	// we create a bucket to store the (worker, bike) pairs
	// for bucket[i] stores the distance=i's pairs
	type pair struct {
		w, b int
	}
	bkt := [2000][]pair{}
	for i := 0; i < nw; i++ {
		for j := 0; j < nb; j++ {
			dis := utils.Abs(workers[i][0]-bikes[j][0]) + utils.Abs(workers[i][1]-bikes[j][1])
			bkt[dis] = append(bkt[dis], pair{w: i, b: j})
		}
	}
	res := make([]int, nw)
	for i := range res {
		res[i] = -1
	}
	set := make(map[int]bool, nb)
	for d := 0; d < 2000; d++ {
		for k := 0; k < len(bkt[d]); k++ {
			if res[bkt[d][k].w] == -1 && !set[bkt[d][k].b] {
				set[bkt[d][k].b] = true
				res[bkt[d][k].w] = bkt[d][k].b
			}
		}
	}
	return res
}
