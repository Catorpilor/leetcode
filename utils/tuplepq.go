package utils

type TuplePriorityQueue [][3]int

func (pq TuplePriorityQueue) Len() int { return len(pq) }

func (pq TuplePriorityQueue) Less(i, j int) bool {
	a, b := pq[i], pq[j]
	if a[0] == b[0] {
		if a[1] == b[1] {
			return a[2] < b[2]
		}
		return a[1] < b[1]
	}
	return a[0] < b[0]
}

func (pq TuplePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq TuplePriorityQueue) Top() [3]int {
	return pq[0]
}

func (pq *TuplePriorityQueue) Push(x interface{}) {
	// n := len(*pq)
	item := x.([3]int)
	*pq = append(*pq, item)
}

func (pq *TuplePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
