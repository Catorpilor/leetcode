package unique

import "math"

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

func backtracking(n int) int {
	if n > 10 {
		return backtracking(10)
	}
	fmax := math.Pow10(n)
	used := make([]bool, 10)
	return permute(float64(0), fmax, &used)
}

func permute(cur, fmax float64, used *[]bool) int {
	if cur >= fmax {
		return 0
	}
	count := 1
	for i := 0; i < 10; i++ {
		if cur == float64(0) && i == 0 {
			// skip the leading 0
			continue
		}
		if !((*used)[i]) {
			(*used)[i] = true
			count += permute(cur*10+float64(i), fmax, used)
			(*used)[i] = false
		}
	}
	return count
}
