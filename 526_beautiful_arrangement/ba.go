package ba

func beautifulArrangement(n int) int {
	var res int
	visited := make([]bool, n+1)
	permute(&res, 1, &visited, n)
	return res
}

func permute(res *int, pos int, visited *[]bool, n int) {
	if pos > n {
		*res++
		return
	}
	for i := 1; i <= n; i++ {
		if !(*visited)[i] && (pos%i == 0 || i%pos == 0) {
			(*visited)[i] = true
			permute(res, pos+1, visited, n)
			(*visited)[i] = false
		}
	}
}

func withMemorization(n int) int {
	curState := make([]byte, n+1)
	for i := range curState {
		curState[i] = 'f'
	}
	cache := make(map[string]int)
	return dfs(cache, &curState, 1)
}

func dfs(cache map[string]int, curState *[]byte, idx int) int {
	if idx == len(*curState) {
		return 1
	}
	ck := string(*curState)
	if v, exists := cache[ck]; exists {
		return v
	}
	var count int
	for i := 1; i < len(*curState); i++ {
		if (*curState)[i] == 'f' && (i%idx == 0 || idx%i == 0) {
			(*curState)[i] = 't'
			count += dfs(cache, curState, idx+1)
			(*curState)[i] = 'f'
		}
	}
	cache[ck] = count
	return count
}
