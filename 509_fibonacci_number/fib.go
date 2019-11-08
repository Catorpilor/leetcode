package fib

func fib(n int) int {
	return iterTopDown(n)
}

func iterTopDown(n int) int {
	// time complexity O(N)
	// space complexity O(1)
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

func matrix(n int) int {
	// time complexity O(lgN)
	// space complexity O(lgN)
	if n <= 1 {
		return n
	}
	matrix := [][]int{[]int{1, 1}, []int{1, 0}}
	matrixPower(matrix, n-1)
	return matrix[0][0]
}

func matrixPower(mat [][]int, n int) {
	if n <= 1 {
		return
	}
	matrixPower(mat, n/2)
	multiply(mat, mat)
	b := [][]int{[]int{1, 1}, []int{1, 0}}
	if n%2 != 0 {
		multiply(mat, b)
	}
}

func multiply(a, b [][]int) {
	x := a[0][0]*b[0][0] + a[0][1]*b[1][0]
	y := a[0][0]*b[0][1] + a[0][1]*b[1][1]
	z := a[1][0]*b[0][0] + a[1][1]*b[1][0]
	w := a[1][0]*b[0][1] + a[1][1]*b[1][1]

	a[0][0] = x
	a[0][1] = y
	a[1][0] = z
	a[1][1] = w
}
