package queue

import "sort"

func reconstruct(people [][]int) [][]int {
	return useInsertionSort(people)
}

// useInsertionSort time complexity O(NlgN), space complexity O(N) or O(1)
func useInsertionSort(people [][]int) [][]int {
	n := len(people)
	// sort people by height in non-increasing order,
	// if height[i] == height[j] then we compare people[i][1] and people[j][1]
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] == people[j][0] && people[i][1] < people[j][1] {
			return true
		} else {
			return false
		}
	})
	// then we pick one people at a time and insert into the res at position people[1]
	res := make([][]int, 0, n)
	for i := range people {
		p := people[i]
		res = append(res[:p[1]], append([][]int{p}, res[p[1]:]...)...)
	}
	return res
}
