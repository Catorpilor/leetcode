package stone

import "sort"

func lastWeightStones(stones []int) int {
	return useSort(stones)
}

// useSort time complexity O(N^2lgN), space complexity O(1)
func useSort(stones []int) int {
	n := len(stones)
	sort.Ints(stones)
	i := n - 2
	for i >= 0 && i+1 < len(stones) {
		if stones[i] == stones[i+1] {
			stones = stones[:i]
			i -= 2
		} else {
			stones[i] = stones[i+1] - stones[i]
			stones = stones[:i+1]
			sort.Ints(stones)
			i--
		}
	}
	if len(stones) > 0 {
		return stones[0]
	}
	return 0
}
