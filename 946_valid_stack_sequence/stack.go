package stack

func isValid(push, pop []int) bool {
	return useGreedy(push, pop)
}

// useGreedy time complexity O(N), space complexity O(N)
func useGreedy(push, pop []int) bool {
	n := len(push)
	st := make([]int, 0, n)
	var j int
	for i := range push {
		st = append(st, push[i])
		nst := len(st)
		for nst > 0 && j < n && st[nst-1] == pop[j] {
			nst--
			st = st[:nst]
			j++
		}
	}
	return j == n
}
