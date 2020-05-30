package reduce

import "sort"

func minSetSize(arr []int) int {
	return useHashmap(arr)
}

// useHashmap time complexity O(N), space compleixty O(N)
func useHashmap(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return n
	}
	set := make(map[int]int)
	want := n >> 1
	for _, c := range arr {
		set[c]++
		if set[c] >= want {
			return 1
		}
	}
	if n == len(set) {
		return want
	}
	unique := make([]int, 0, len(set))
	for k := range set {
		unique = append(unique, k)
	}
	sort.Slice(unique, func(i, j int) bool {
		return set[unique[i]] > set[unique[j]]
	})
	var count, ans int
	for i := range unique {
		count += set[unique[i]]
		if count >= want {
			ans = i + 1
			break
		}
	}
	return ans
}
