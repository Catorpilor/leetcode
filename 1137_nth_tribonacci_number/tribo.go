package tribo

func tribonacci(n int) int {
	x0, x1, x2 := 0, 1, 1
	if n <= 1 {
		return n
	}
	var tmp int
	for i := 3; i <= n; i++ {
		x2 += x0 + x1
		tmp = x1
		x1 = x2 - x0 - x1
		x0 = tmp
	}
	return x2
}

// time complexity O(38) = O(1)
// space complexity O(38)
func precompute(res *[]int, n int) {
	for i := 3; i < n; i++ {
		(*res)[i] = (*res)[i-2] + (*res)[i-3] + (*res)[i-1]
	}
}

func memorization(n int) int {
	if n <= 1 {
		return n
	}
	res := make([]int, 38)
	res[1], res[2] = 1, 1
	precompute(&res, 38)
	return res[n]
}
