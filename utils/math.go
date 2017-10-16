package utils

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var EPSILON float64 = 0.00001

func FloatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func Maxf(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	if a == 0 {
		return 0
	}
	return a
}
