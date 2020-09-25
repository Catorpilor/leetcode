package pooling

func carPooling(trips [][]int, capacity int) bool {
	return useBucket(trips, capacity)
}

// useBucket time complexity O(N+1001), space complexity O(1)
func useBucket(trips [][]int, capacity int) bool {
	stops := [1001]int{}
	for _, trip := range trips {
		p, s, e := trip[0], trip[1], trip[2]
		stops[s] += p
		stops[e] -= p
	}
	for i := 0; i < 1001 && capacity > 0; i++ {
		capacity -= stops[i]
	}
	return capacity >= 0
}
