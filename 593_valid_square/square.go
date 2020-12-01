package square

func isValid(p1, p2, p3, p4 []int) bool {
	return useMath(p1, p2, p3, p4)
}

// useMath time complexity O(1), space complexity O(1)
func useMath(p1, p2, p3, p4 []int) bool {
	dist := func(p1, p2 []int) int {
		return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
	}
	ls := []int{dist(p1, p2), dist(p1, p3), dist(p1, p4), dist(p2, p3), dist(p2, p4), dist(p3, p4)}
	set := make(map[int]bool, 6)
	for _, l := range ls {
		set[l] = true
	}
	return !set[0] && len(set) == 2
}
