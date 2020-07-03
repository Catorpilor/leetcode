package mvoe

import "math"

func findMaxValueOfEquation(points [][]int, k int) int {
	return useBruteForce(points, k)
}

// useBruteForce time complexity O(N^2), space complexity O(N)
// got tle
func useBruteForce(points [][]int, k int) int {
	type pair struct {
		x, y int
	}
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
