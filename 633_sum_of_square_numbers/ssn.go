package ssn

func JudgeSSN(num int) bool {
	// determine boundray
	// caculate root of num
	v := num
	for v*v > num {
		v = ((num / v) + v) >> 1
	}
	for a, b := 0, v; a <= b; {
		if a*a+b*b < num {
			a += 1
		} else if a*a+b*b > num {
			b -= 1
		} else {
			return true
		}
	}
	return false
}
