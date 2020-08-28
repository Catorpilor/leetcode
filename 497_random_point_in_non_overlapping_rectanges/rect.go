package rect

import "math/rand"

// Solution is the base struct.
type Solution struct {
	rects       [][]int
	pSumOfAreas []int
	n           int
}

// Constructor creates a solution based on the rects
func Constructor(rects [][]int) *Solution {
	n := len(rects)
	sum := 0
	prefix := make([]int, 0, n)
	for _, rect := range rects {
		sum += (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1)
		prefix = append(prefix, sum)
	}
	return &Solution{rects: rects, pSumOfAreas: prefix, n: n}
}

// Pick pick a random position based on the weight of the rectangles
func (s *Solution) Pick() []int {
	target := rand.Intn(s.pSumOfAreas[s.n-1]) + 1
	l, r := 0, s.n-1
	for l <= r {
		mid := l + (r-l)/2
		if s.pSumOfAreas[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	bound := s.rects[l]
	ans := make([]int, 2)
	ans[0] = bound[0] + rand.Intn(bound[2]-bound[0]+1)
	ans[1] = bound[1] + rand.Intn(bound[3]-bound[1]+1)
	return ans
}
