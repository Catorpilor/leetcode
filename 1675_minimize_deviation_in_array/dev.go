package dev

import (
	"container/heap"
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func minimumDeviation(nums []int) int {
	return usePriorityQueue(nums)
}

type MaxPQ []int

func (pq MaxPQ) Len() int { return len(pq) }

func (pq MaxPQ) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq MaxPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq MaxPQ) Top() int {
	return pq[0]
}

func (pq *MaxPQ) Push(x interface{}) {
	// n := len(*pq)
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *MaxPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// usePriorityQueue time complexity O(N*log(N)), space complexity O(N)
func usePriorityQueue(nums []int) int {
	var pq MaxPQ
	minV := math.MaxInt32
	for _, num := range nums {
		if num&1 != 0 {
			num <<= 1
		}
		heap.Push(&pq, num)
		minV = utils.Min(minV, num)
	}
	res := math.MaxInt32
	for pq.Top()%2 == 0 {
		res = utils.Min(res, pq.Top()-minV)
		minV = utils.Min(minV, pq.Top()>>1)
		heap.Push(&pq, pq.Top()>>1)
		heap.Pop(&pq)
	}
	return utils.Min(res, pq.Top()-minV)
}
