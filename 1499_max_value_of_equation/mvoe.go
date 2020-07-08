package mvoe

import (
	"math"

	"github.com/catorpilor/leetCode/utils"
)

func findMaxValueOfEquation(points [][]int, k int) int {
	return useBruteForce(points, k)
	// return useDeque(points, k)
}

type pair struct {
	x, y int
}

// useBruteForce time complexity O(N^2), space complexity O(N)
// got tle
func useBruteForce(points [][]int, k int) int {
	n := len(points)
	pairs := make([]pair, 0, n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if points[j][0]-points[i][0] <= k {
				pairs = append(pairs, pair{i, j})
			} else {
				break
			}
		}
	}
	ans := math.MinInt32
	for _, p := range pairs {
		i, j := p.x, p.y
		tmp := points[i][1] + points[j][1] + points[j][0] - points[i][0]
		if tmp > ans {
			ans = tmp
		}
	}
	return ans
}

// useDeque time complexity O(N), space complexity O(N)
func useDeque(points [][]int, k int) int {
	// deq is a monotonic decreasing stack, dep.Front() is the largest one.
	deq := utils.NewDeque()
	res := math.MinInt32
	for _, point := range points {
		x, y := point[0], point[1]
		for !deq.IsEmpty() && deq.Front().(pair).y < x-k {
			deq.PopFront()
		}
		if !deq.IsEmpty() {
			res = utils.Max(res, deq.Front().(pair).x+x+y)
		}
		for !deq.IsEmpty() && deq.Back().(pair).x <= y-x {
			deq.PopBack()
		}
		deq.PushBack(pair{y - x, x})
	}
	return res
}
