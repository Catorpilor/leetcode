package path

import (
	"container/heap"
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func minimumEffort(heights [][]int) int {
	return useDij(heights)
}

func useDFS(heights [][]int) int {
	m := len(heights)
	if m < 1 {
		return 0
	}
	n := len(heights[0])
	if n < 1 {
		return 0
	}
	visited := make([]bool, m*n)
	ans := math.MaxInt32
	helper(heights, &ans, 0, 0, m, n, visited, 0, 0)
	return ans
}

func helper(heights [][]int, ans *int, i, j, m, n int, visited []bool, prev, curDiff int) {
	pos := i*m + j
	// fmt.Printf("i(%d)>=m(%d) ?%t, i<0?%t, j<0?%t, j(%d)>=n(%d)?%t\n", i, m, i >= m, i < 0, j < 0, j, n, j >= n)
	if i < 0 || i >= m || j < 0 || j >= n || visited[pos] {
		// fmt.Printf("(%d, %d) out\n", i, j)
		return
	}
	diff := utils.Abs(prev - heights[i][j])
	if diff > curDiff {
		curDiff = diff
	}
	// fmt.Printf("curDiff: %d\n", curDiff)
	if i == m-1 && j == n-1 {
		// update ans if necessary
		*ans = utils.Min(*ans, curDiff)
		return
	}
	visited[pos] = true
	dirs := [5]int{-1, 0, 1, 0, -1}
	for k := 0; k < 4; k++ {
		// fmt.Printf("cur i, j (%d, %d), m=%d, n=%d\n", i, j, m, n)
		helper(heights, ans, i+dirs[k], j+dirs[k+1], m, n, visited, heights[i][j], curDiff)
	}
	visited[pos] = false
}

type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() []int        { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func useDij(heights [][]int) int {
	m := len(heights)
	if m < 1 {
		return 0
	}
	n := len(heights[0])
	if n < 1 {
		return 0
	}
	visited := make([]bool, m*n)
	mh := &MinHeap{}
	heap.Init(mh)
	heap.Push(mh, []int{0, 0, 0})
	dirs := [5]int{-1, 0, 1, 0, -1}
	for mh.Len() > 0 {
		tp := heap.Pop(mh).([]int)
		pos := tp[0]*n + tp[1]
		if visited[pos] {
			continue
		}
		visited[pos] = true
		if tp[0] == m-1 && tp[1] == n-1 {
			return tp[2]
		}
		for k := 0; k < 4; k++ {
			nx, ny := tp[0]+dirs[k], tp[1]+dirs[k+1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && !visited[nx*n+ny] {
				heap.Push(mh, []int{nx, ny, utils.Max(tp[2], utils.Abs(heights[nx][ny]-heights[tp[0]][tp[1]]))})
			}
		}
	}
	return 0
}
