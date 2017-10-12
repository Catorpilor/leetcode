package alter

func HasAlter(n int) bool {
	prev := n & 1
	n >>= 1
	for n > 0 {
		if (n&1)^prev == 0 {
			return false
		}
		prev = n & 1
		n >>= 1
	}
	return true
}

func HasAlter2(n int) bool {
	// if n has interleaving bits
	// then (n>>1) + n must be all 1s
	// eg 5 -> 101 5>>1 = 10 + 101 = 111
	// then for every n > 0 we count how many ones and check if
	// it equals to its bit length
	np := n + (n >> 1)
	bitLength, numOfOne := 0, 0
	for np > 0 {
		numOfOne += (np & 1)
		bitLength += 1
		np >>= 1
	}
	return bitLength == numOfOne
}
