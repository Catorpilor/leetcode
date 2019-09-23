package nge

import "fmt"

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
	// find the largest digit in digits[i+1:] that is smaller than digits[i]
	curMaxIdx := i + 1
	for j := i + 2; j < n; j++ {
		if digits[j] < digits[i] && digits[j] > digits[curMaxIdx] {
			curMaxIdx = j
		}
	}
	// swap i with curMaxIdx
	digits[i], digits[curMaxIdx] = digits[curMaxIdx], digits[i]
	for l, r := 0, curMaxIdx-1; l < r; l, r = l+1, r-1 {
		digits[l], digits[r] = digits[r], digits[l]
	}
	fmt.Printf("input %d final digits: %v\n", num, digits)
	return -1
}
