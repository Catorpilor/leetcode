package fib

func fib(n int) int {
	x0, x1 := 0, 1
	if n == 0 {
		return x0
	}
	for i := 2; i <= n; i++ {
		x1 += x0
		x0 = x1 - x0
	}
	return x1
}
