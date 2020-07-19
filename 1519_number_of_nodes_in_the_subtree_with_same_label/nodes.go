package nodes

func numberOfNodes(n int, edges [][]int, labels string) []int {
	return useDFS(n, edges, labels)
}

var adj map[int][]int
var c [26]int
var ans []int

func useDFS(n int, edges [][]int, labels string) []int {
	adj = make(map[int][]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	ans = make([]int, n)
	helper(0, -1, labels)
	return ans
}

func helper(u, p int, labels string) {
	c[labels[u]-'a']++
	prev := c[labels[u]-'a']
	for _, v := range adj[u] {
		if v^p != 0 {
			helper(v, u, labels)
		}
	}
	ans[u] = c[labels[u]-'a'] - prev + 1
}
