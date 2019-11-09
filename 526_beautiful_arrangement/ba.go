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
