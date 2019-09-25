package course

func topologicalSort(numCourses int, prerequisites [][]int) []int {
	n := len(prerequisites)
	res := make([]int, 0, numCourses)
	if n == 0 {
		for i := 0; i < numCourses; i++ {
			res = append(res, i)
		}
		return res
	}
	g := make(map[int][]int, numCourses)
	for _, v := range prerequisites {
		g[v[1]] = append(g[v[1]], v[0])
	}
	// count indegress of each nodes
	degs := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		for _, v := range g[i] {
			degs[v]++
		}
	}
	for i := range degs {
		if degs[i] == 0 {
			res = append(res, i)
		}
	}

	for i := 0; i < len(res); i++ {
		for _, v := range g[res[i]] {
			degs[v]--
			if degs[v] == 0 {
				res = append(res, v)
			}
		}
	}

	if len(res) != numCourses {
		return []int{}
	}
	return res
}
