package course

func CanFinish(numCourses int, pres [][]int) bool {
	n := len(pres)
	if n == 0 {
		return true
	}
	visited, onpath := make([]bool, numCourses), make([]bool, numCourses)
	g := make(map[int][]int)
	// fill g with pres
	for _, v := range pres {
		g[v[1]] = append(g[v[1]], v[0])
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] {
			continue
		}
		if dfsCycle(g, &visited, &onpath, i) {
			return false
		}
	}
	return true
}

func dfsCycle(g map[int][]int, visited, onpath *[]bool, node int) bool {
	// prune the duplicated check
	if (*visited)[node] {
		return false
	}
	// mark node as visited
	(*visited)[node], (*onpath)[node] = true, true

	for _, adj := range g[node] {
		if (*onpath)[adj] || dfsCycle(g, visited, onpath, adj) {
			return true
		}
	}
	// make node onpath to false
	(*onpath)[node] = false
	return false
}
