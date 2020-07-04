package cell

import (
	"strconv"
	"strings"
)

func afterNDays(cell []int, n int) []int {
	// return useBruteForce(cell, n)
	return useHashmap(cell, n)
}

// time complexity O(N), space complexity O(1)
func useBruteForce(cell []int, n int) []int {
	for i := 1; i <= n; i++ {
		cell = nextDay(cell)
		// fmt.Println(cell)
	}
	return cell
}

func nextDay(cell []int) []int {
	n := len(cell)
	tmp := make([]int, 8)
	// copy(tmp, cell)
	for i := 1; i < n-1; i++ {
		if cell[i-1]^cell[i+1] == 1 {
			tmp[i] = 0
		} else {
			tmp[i] = 1
		}
	}
	return tmp
}

func useHashmap(cell []int, n int) []int {
	orignal := make([]int, 8)
	copy(orignal, cell)
	set := make(map[string]bool)
	hasCycle := false
	var length int
	for i := 0; i < n; i++ {
		nd := nextDay(cell)
		key := intSliceToString(nd)
		if set[key] {
			// we found a cycle
			hasCycle = true
			break
		}
		set[key] = true
		length++
		cell = nd
	}
	if hasCycle {
		n %= length
		for i := 0; i < n; i++ {
			cell = nextDay(cell)
		}
	}
	return cell

}

func intSliceToString(a []int) string {
	b := make([]string, 8)
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	return strings.Join(b, "")
}
