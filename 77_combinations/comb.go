package comb

func combinations(n, k int) [][]int {
	var res [][]int
	if n < k {
		return res
	}

	if n == k {
		tmp := make([]int, 0, n)
		for i := 1; i <= n; i++ {
			tmp = append(tmp, i)
		}
		res = append(res, tmp)
	} else {
		permute(&res, 1, n, k, []int{})
	}
	return res
}

func permute(res *[][]int, pos, n, k int, bid []int) {
	if len(bid) == k {
		tmp := make([]int, k)
		copy(tmp, bid)
		*res = append(*res, tmp)
		return
	}
	for i := pos; i <= n; i++ {
		bid = append(bid, i)
		permute(res, i+1, n, k, bid)
		bid = bid[:len(bid)-1]
	}
}
