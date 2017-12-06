package rect

import (
	"github.com/catorpilor/leetcode/utils"
)

func LargestRectangle(heights []int) int {
	n := len(heights)
	if n < 1 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}
	var top, i, areaSoFar, maxArea int
	s := utils.NewStack()
	for i < n {
		// if stack is empty or heights[i] >= heights[top] we keep pushing
		if s.IsEmpty() || heights[s.Top().(int)] <= heights[i] {
			s.Push(i)
			i++
		} else {
			// heights[i] is lower than heights[top]
			// pop
			v := s.Pop().(int)
			if s.IsEmpty() {
				areaSoFar = heights[v] * i
			} else {
				top = s.Top().(int)
				areaSoFar = heights[v] * (i - top - 1)
			}
			if areaSoFar > maxArea {
				maxArea = areaSoFar
			}
		}
	}
	for !s.IsEmpty() {
		v := s.Pop().(int)
		if s.IsEmpty() {
			areaSoFar = heights[v] * i
		} else {
			top = s.Top().(int)
			areaSoFar = heights[v] * (i - top - 1)
		}
		if areaSoFar > maxArea {
			maxArea = areaSoFar
		}

	}
	return maxArea
}

func LargestRectangle2(heights []int) int {
	// brute force
	n := len(heights)
	if n < 1 {
		return 0
	}
	var minHeight, areaSofar, maxArea int
	for i := 0; i < n; i++ {
		minHeight = heights[i]
		if minHeight*1 > maxArea {
			maxArea = minHeight
		}
		for j := i + 1; j < n; j++ {
			if heights[j] < minHeight {
				minHeight = heights[j]
			}
			areaSofar = minHeight * (j - i + 1)
			if areaSofar > maxArea {
				maxArea = areaSofar
			}
		}
	}
	return maxArea
}
