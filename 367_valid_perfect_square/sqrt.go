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

func useMath(num int) bool {
	i := 1
	for num > 0 {
		num, i = num-i, i+2
	}
	return num == 0
}

func useBinarySearch(num int) bool {
	// binary search
	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}
	l, r := 1, num/2
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
