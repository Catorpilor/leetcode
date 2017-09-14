package happy

func squareDigit(n int) int {
	ret := 0
	for n > 0 {
		b := n % 10
		ret += b * b
		n = n / 10
	}
	return ret
}

func IsHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = squareDigit(slow)
		fast = squareDigit(fast)
		fast = squareDigit(fast)
		if slow == 1 {
			return true
		}
		if slow == fast {
			break
		}
	}
	return false
}
