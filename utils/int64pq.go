package utils

type Int64PriorityQueue []int64

func (pq Int64PriorityQueue) Len() int { return len(pq) }

func (pq Int64PriorityQueue) Less(i, j int) bool {
    return pq[i] < pq[j]
}

func (pq Int64PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq Int64PriorityQueue) Top() int64 {
    return pq[0]
}

func (pq *Int64PriorityQueue) Push(x interface{}) {
    // n := len(*pq)
    item := x.(int64)
    *pq = append(*pq, item)
}

func (pq *Int64PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}
