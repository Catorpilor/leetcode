package verticies

func minimumVerticies(n int, edge [][]int) []int {
	return useIndegree(n, edge)
}

// useIndegree time complexity O(E), space complexity O(N)
func useIndegree(n int, edges [][]int) []int {
	indegree := make([]int, n)
	for _, edge := range edges {
		indegree[edge[1]]++
	}
	ans := make([]int, 0, n)
	for i := range indegree {
		if indegree[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}
