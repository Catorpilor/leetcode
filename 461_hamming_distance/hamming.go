package hamming

func Distance(x, y int) int {
	z := x ^ y
	var ret int
	for z > 0 {
		ret += 1
		z &= (z - 1)
	}
	return ret
}
