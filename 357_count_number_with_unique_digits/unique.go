package unique

func countUniqueNumbers(n int) int {
	return dp(n)
}

// based on discussion
// f(1) = 10  // (0..9) unique numbers
// f(2) = 9*9 // choose i from (1..9) choose j from (0..9) except i
// f(3) = f(2) * 8 // the third postion just have 8 choices
// f(4) = f(3) * 7
// ...
// f(10) = f(9) * 1
// f(11) = f(12) = 0
func dp(n int) int {
	if n == 0 {
		return 1
	}
	res := 10
	uniqueDigits := 9
	avaliableNumber := 9
	for n > 1 && avaliableNumber > 0 {
		uniqueDigits *= avaliableNumber
		res += uniqueDigits
		avaliableNumber--
		n--
	}
	return res
}
