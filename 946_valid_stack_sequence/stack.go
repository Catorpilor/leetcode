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

// useConstantSpace time complexity O(N), space complexity O(1)
func useConstantSpace(pushed, popped []int) bool {
	n := len(pushed)
	k := -1 // -1 means empty stack
	var j int
	for i := range pushed {
		k++
		pushed[k] = pushed[i]
		for k != -1 && pushed[k] == popped[i] {
			k--
			i++
		}
	}
	return k == -1
}
