package matrics

func countMatrics(mat [][]int) int {
	return useOneRowHelper(mat)
}

// useOneRowHelper time complexity O(MNM), space complexity O(N)
func useOneRowHelper(mat [][]int) int {
	m := len(mat)
	if m < 1 {
		return 0
	}
	n := len(mat[0])
	if n < 1 {
		return 0
	}
	var ans int
	for up := 0; up < m; up++ {
		local := make([]int, n)
		for i := range local {
			local[i] = 1
		}
		for down := up; down < m; down++ {
			for k := 0; k < n; k++ {
				local[k] &= mat[down][k]
			}
			ans += countOneRow(local)
		}
	}
	return ans
}

func countOneRow(arr []int) int {
	n := len(arr)
	var ans, length int

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			length = 0
		} else {
			length++
		}
		ans += length
	}
	return ans
}
