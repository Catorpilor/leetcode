package rect

func largestRectangleArea(heights []int) int {
	return useStack(heights)
}

// useStack time compelxity O(N), space complexity O(N)
func useStack(heights []int) int {
	n := len(heights)
	if n < 1 {
		return 0
	}
	st := make([]int, 0, n)
	var i, areaSoFar, maxArea int

	for i < n {
		nst := len(st)
		if nst == 0 || heights[st[nst-1]] <= heights[i] {
			st = append(st, i)
			i++
		} else {
			// heights[i] < heights[st.top()]
			v := st[nst-1]
			nst--
			st = st[:nst]
			if nst == 0 {
				areaSoFar = heights[v] * i
			} else {
				areaSoFar = heights[v] * (i - st[nst-1] - 1)
			}
			if areaSoFar > maxArea {
				maxArea = areaSoFar
			}
		}
	}
	// monotonically increasing heights
	nst := len(st)
	for nst > 0 {
		v := st[nst-1]
		nst--
		st = st[:nst]
		if nst == 0 {
			areaSoFar = heights[v] * i
		} else {
			areaSoFar = heights[v] * (i - st[nst-1] - 1)
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
