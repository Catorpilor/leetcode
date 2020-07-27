package wb

func watterBottles(numBottles, numExchange int) int {
	return useMath(numBottles, numExchange)
}

// useMath time complexity O(1), space compelxity O(1)
func useMath(numBottles, numExchange int) int {
	ans := numBottles
	var rem int
	for numBottles >= numExchange {
		rem = numBottles % numExchange
		numBottles /= numExchange
		ans += numBottles
		numBottles += rem
	}
	return ans
}
