package values

import "sort"

func getStrongest(arr []int, k int) []int {
	return useSort(arr, k)
}

// useSort time complexity O(NlgN), space complexity O(N)
func useSort(arr []int, k int) []int {
	n := len(arr)
	if n <= 1 || k == n {
		return arr
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	// find the median
	median := arr[(n-1)/2]
	// create a temp arr
	tmp := make([]int, n)
	copy(tmp, arr)
	sort.Slice(tmp, func(i, j int) bool {
		if weight(tmp[i], median) > weight(tmp[j], median) {
			return true
		} else if weight(tmp[i], median) == weight(tmp[j], median) && tmp[i] > tmp[j] {
			return true
		} else {
			return false
		}
	})
	return tmp[:k]
}

func weight(num, m int) int {
	return abs(num - m)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
