package course

import (
	"container/heap"
	"fmt"
)

func minSems(n int, deps [][]int, k int) int {
	return useTopSort(n, deps, k)
}

// useTopSort
func useTopSort(n int, dependencies [][]int, k int) int {
	nd := len(dependencies)
	if nd == 0 {
		return n/k + n%k
	}
	adj := make([][]int, n+1)
	indgree := make([]int, n+1)
	// we need to calculagte the out degrees
	for _, dep := range dependencies {
		p, q := dep[0], dep[1]
		adj[p] = append(adj[p], q)
		indgree[q]++
	}
	visited := make([]bool, n+1)
	var ans, count int
	// find ingree == 0
	for {
		starts := make([]int, 0, n+1)
		for i := 1; i <= n; i++ {
			if indgree[i] == 0 && !visited[i] {
				starts = append(starts, i)
				visited[i] = true
				count++
			}
		}
		fmt.Printf("pick starts: %v, count: %d\n", starts, count)
		// fmt.Println(starts)
		// this logic is wrong.....
		ns := len(starts)
		if ns <= k {
			ans++
		} else {
			ans += ns/k + ns%k
		}
		for i := range starts {
			child := adj[starts[i]]
			for j := range child {
				indgree[child[j]]--
			}
		}
		if count == n {
			break
		}
	}
	return ans
}

func useTopSortWithPriorityQueue(n int, dependencies [][]int, k int) int {
	nd := len(dependencies)
	if nd == 0 {
		return n/k + 1
	}
	adj := make([][]int, n+1)
	indegrees := make([]int, n+1)
	// this not working... we should use depth instead
	// failed example n=12, deps{[1,2],[1,3],[7,5],[7,6],[4,8],[8,9],[9,10],[10,11],[11,12]}, k=2
	// wanted 6 but use out degrees we got 7.
	outdegrees := make([]int, n+1)
	for _, dep := range dependencies {
		p, q := dep[0], dep[1]
		adj[p] = append(adj[p], q)
		indegrees[q]++
		outdegrees[p]++
	}
	// find node's indegree eq 0
	batch := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		if indegrees[i] == 0 {
			batch = append(batch, i)
		}
	}
	pq := make(PriorityQueue, 0, n+1)
	for _, node := range batch {
		pq = append(pq, &item{outdegree: outdegrees[node], index: node})
	}
	heap.Init(&pq)
	fmt.Printf("before, pq(%d)\n", pq.Len())
	var ans, take int
	for pq.Len() > 0 {
		next := make([]*item, 0, n)
		for take < k && pq.Len() > 0 {
			l := heap.Pop(&pq).(*item)
			fmt.Printf("pop item{outd: %d, node:%d} out of pq(%d)\n", l.outdegree, l.index, pq.Len())
			for _, node := range adj[l.index] {
				indegrees[node]--
				if indegrees[node] == 0 {
					next = append(next, &item{index: node, outdegree: outdegrees[node]})
					fmt.Printf("append item(index:%d, outdegree:%d)\n", node, outdegrees[node])
				}
			}
			take++
		}
		take = 0
		ans++
		fmt.Printf("push next len:%d to pq\n", len(next))
		for _, c := range next {
			heap.Push(&pq, c)
		}
	}
	return ans
}

type item struct {
	outdegree int
	index     int
}

// PriorityQueue is a priority queue based on the outdegree
type PriorityQueue []*item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// fmt.Printf("pq: %v\n", pq)
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].outdegree > pq[j].outdegree
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) Top() int {
	return pq[0].index
}
