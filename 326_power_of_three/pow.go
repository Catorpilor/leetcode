package pow

func Pow(n int) bool {
	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n = n / 3
	}
	return n == 1
}

func Pow3(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n = n / 3
	}
	return n == 1
}
