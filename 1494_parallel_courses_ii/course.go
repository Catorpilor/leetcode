package course

import (
	"container/heap"
	"fmt"
	"math/bits"

	"github.com/catorpilor/leetcode/utils"
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
	// update: depth approach is not working either, the only way to solve it is to use dp....
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

// useBitMaskDP time complexity O(3^N), space complexity O(N)
func useBitMaskDP(n int, dependencies [][]int, k int) int {
	nd := len(dependencies)
	if nd == 0 {
		return n/k + 1
	}
	pre := make([]int, n)
	for _, edge := range dependencies {
		p, q := edge[0]-1, edge[1]-1
		pre[q] |= 1 << p
	}
	dp := make([]int, 1<<n)
	for i := 1; i < len(dp); i++ {
		dp[i] = n
	}
	for i := 0; i < (1 << n); i++ {
		// i is the bit representation of a combination of courses.
		// dp[i] is the minimum days to complete all the courses
		var ex int
		// find  out ex, the bit representation of the all courses that we can study now
		// since we have i and pre[j], we know course j can be studied if i contains all it's prerequisites ((i & pre[j]) == pre[j])
		for j := 0; j < n; j++ {
			if i&pre[j] == pre[j] {
				ex |= 1 << j
			}
		}
		ex &= ^i
		// enumerate all the bit 1 combinations of ex
		// this is a typical method to enumerate all subsets of a bit representation:
		// for (int i = s; i; i = (i - 1) &ｓ)
		for s := ex; s != 0; s = (s - 1) & ex {
			if bits.OnesCount(uint(s)) <= k {
				dp[i|s] = utils.Min(dp[i|s], dp[i]+1)
			}
		}

	}
	return dp[1<<n-1]
}
