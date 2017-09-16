package primes

import (
	"math"
)

func CountPrimes(n int) int {
	nums := make([]bool, n)
	var ret int
	for i := 2; i < n; i++ {
		if nums[i] == false {
			ret += 1
			for j := 2; i*j < n; j++ {
				nums[i*j] = true
			}
		}
	}
	return ret
}

func CountPrimes2(n int) int {
	nums := make([]bool, n)
	nums[0], nums[1] = true, true
	sqn := math.Sqrt(float64(n))
	for i := 0; i <= int(sqn); i++ {
		if nums[i] == false {
			for j := i * i; j < n; j += i {
				nums[j] = true
			}
		}
	}
	var ret int
	for i := range nums {
		if nums[i] == false {
			ret += 1
		}
	}
	return ret
}
