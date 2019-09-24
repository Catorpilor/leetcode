package nge

import (
	"math"
)

func nextGreaterElement(num int) int {
	var digits []int
	temp := num
	for temp > 0 {
		d := temp % 10
		digits = append(digits, d)
		temp /= 10
	}
	n := len(digits)
	if n <= 1 {
		return -1
	}
	var i int
	for ; i < n-1; i++ {
		if digits[i] > digits[i+1] {
			break
		}
	}
	if i == n-1 {
		// num in descresing order
		return -1
	}
	// find the first digit in digits[:i] that is smaller than or equal to digits[i]
	// larger than digits[i+1]
	x, smallest := digits[i+1], i
	for j := 0; j < i; j++ {
		if digits[j] > x && digits[j] <= digits[smallest] {
			// we are in reverse order
			// just find the first digit
			smallest = j
			break
		}
	}
	// swap i with curMaxIdx
	digits[i+1], digits[smallest] = digits[smallest], digits[i+1]
	// fmt.Printf("swaped digs: %v\n", digits)
	for l, r := 0, i; l < r; l, r = l+1, r-1 {
		digits[l], digits[r] = digits[r], digits[l]
	}
	// fmt.Printf("input %d final digits: %v\n", num, digits)
	var res int64
	// res = (int64)(digits[0])
	k := 1
	for p := 0; p < n; p++ {
		res += (int64)(digits[p] * k)
		k *= 10
	}
	// fmt.Printf("res is %d\n", res)
	if res > math.MaxInt32 {
		return -1
	}

	return int(res)
}
