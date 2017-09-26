package sqrt

func IsPerfectSqure(num int) bool {
	v := num
	for v*v > num {
		v = (num/v + v) >> 1
	}
	return v*v == num
}

// 1 = 1
// 4 = 1 + 3
// 9 = 1 + 3 + 5
// 16 = 1 + 3 + 5 + 7
// 25 = 1 + 3 + 5 + 7 + 9
// n*n = 1 + 3 + 5 + 7 + 9 + ... (n * (1 + 2n-1) / 2 )

func IsPerfectSqure2(num int) bool {
	i := 1
	for num > 0 {
		num, i = num-i, i+2
	}
	return num == 0
}

func IsPerfectSqure3(num int) bool {
	// binary search
	if num <= 0 {
		return false
	}
	l, r := 1, num
	for l < r {
		mid := l + (r-l)/2 + 1
		if num/mid >= mid {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l*l == num
}
