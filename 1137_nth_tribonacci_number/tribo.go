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
