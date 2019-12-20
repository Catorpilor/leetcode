package utils

func NumOfOnes(a int) int {
	var res int
	for a > 0 {
		a &= (a - 1)
		res++
	}
	return res
}
