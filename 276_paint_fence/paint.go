package paint

func NumWays(n, k int) int {
	if n == 0 || k == 0 || (k == 1 && n > 2) {
		return 0
	}
	a := k     // single fence
	b := k * k // two fences
	if n == 1 {
		return a
	}
	if n == 2 {
		return b
	}
	// equation w(n) = (k-1) * (w(n-1)+w(n-2))
	// two cases
	// a. p(n) == p(n-1) => p(n-2) != p(n-1) => (k-1)*w(n-2)
	// b. p(n-1) != p(n) => w(n) = (K-1)*w(n-1)
	for i := 3; i <= n; i++ {
		c := (k - 1) * (a + b)
		a = b
		b = c
	}
	return b
}
