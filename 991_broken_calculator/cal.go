package cal

func brokenCalc(x, y int) int {
	return useRec(x, y)
}

// useRec time complexity O(log(Y)), space complexity O(logY)
func useRec(x, y int) int {
	if x >= y {
		return x - y
	}
	if y&1 == 1 {
		return 1 + useRec(x, y+1)
	}
	return 1 + useRec(x, y<<1)
}
