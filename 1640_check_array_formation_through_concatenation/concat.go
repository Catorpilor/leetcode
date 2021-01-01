package concat

func canFormArray(arr []int, pcs [][]int) bool {
	return useHashmap(arr, pcs)
}

// useHashmap time complexity O(N) where (N is the length of arr), space compelxity O(N)
func useHashmap(arr []int, pcs [][]int) bool {
	n := len(arr)
	set := make(map[int]int, n)
	for i := range arr {
		set[arr[i]] = i
	}
	for i := range pcs {
		cur := pcs[i]
		if len(cur) == 1 {
			if _, exists := set[cur[0]]; !exists {
				return false
			}
			continue
		}
		start := -1
		cj := 0
		for j := range cur {
			if p, exists := set[cur[j]]; exists {
				if start == -1 {
					start = p
				} else {
					if j-cj != p-start {
						return false
					}
				}
			} else {
				return false
			}
		}
	}
	return true
}
