package cell

func afterNDays(cell []int, n int) []int {
	return useBruteForce(cell, n)
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
