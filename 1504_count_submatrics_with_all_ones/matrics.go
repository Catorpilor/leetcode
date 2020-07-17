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

// useStack time complexity O(MN), space complexity O(N)
func useStack(mat [][]int) int {
	m := len(mat)
	if m < 1 {
		return 0
	}
	n := len(mat[0])
	if n < 1 {
		return 0
	}
	var res int
	h := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != 0 {
				h[j] = h[j] + 1
			} else {
				h[j] = 0
			}
		}
		res += helper(h)
	}
	return res
}

func helper(arr []int) int {
	n := len(arr)
	local := make([]int, n)
	st := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nst := len(st)
		for nst > 0 && arr[st[nst-1]] >= arr[i] {
			nst--
			st = st[:nst]
		}
		if nst > 0 {
			preIndex := st[nst-1]
			local[i] = arr[preIndex]
			local[i] += arr[i] * (i - preIndex)
		} else {
			local[i] = arr[i] * (i + 1)
		}
		st = append(st, i)
	}
	var res int
	for _, c := range local {
		res += c
	}
	return res
}
