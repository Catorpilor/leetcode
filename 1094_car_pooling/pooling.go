package pooling

import (
	"container/heap"
	"sort"
)

func carPooling(trips [][]int, capacity int) bool {
	// return useBucket(trips, capacity)
	return usePriorityQueue(trips, capacity)
}

// useBucket time complexity O(N+1001), space complexity O(1)
func useBucket(trips [][]int, capacity int) bool {
	stops := [1001]int{}
	for _, trip := range trips {
		p, s, e := trip[0], trip[1], trip[2]
		stops[s] += p
		stops[e] -= p
	}
	for i := 0; i < 1001 && capacity > 0; i++ {
		capacity -= stops[i]
	}
	return capacity >= 0
}

// usePriorityQueue time complexity O(nLgN) space complexity  O(N)
func usePriorityQueue(trips [][]int, capacity int) bool {
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})
	pq := make(lpq, 0, len(trips))
	for _, trip := range trips {
		tn, s := trip[0], trip[1]
		for pq.Len() > 0 && pq.Top().([]int)[2] <= s {
			item := heap.Pop(&pq).([]int)
			capacity += item[0]
		}
		capacity -= tn
		if capacity < 0 {
			return false
		}
		heap.Push(&pq, trip)
	}
	return true
}

type lpq [][]int

func (pq lpq) Len() int { return len(pq) }

func (pq lpq) Less(i, j int) bool {
	return pq[i][2] < pq[j][2]
}

func (pq lpq) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *lpq) Push(x interface{}) {
	item := x.([]int)
	*pq = append((*pq), item)
}

func (pq *lpq) Top() interface{} {
	// n := len(*pq)
	return (*pq)[0]
}

func (pq *lpq) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[0]
	old[0] = nil
	*pq = old[1:n]
	return item
}
