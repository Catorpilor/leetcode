package randweight

import "math/rand"

// RW is a random weight collection
type RW struct {
	accSum []int
	n      int
}

// Constructor creates a RW based on arr
func Constructor(arr []int) *RW {
	n := len(arr)
	tmp := make([]int, n)
	copy(tmp, arr)
	for i := 1; i < n; i++ {
		tmp[i] += tmp[i-1]
	}
	return &RW{accSum: tmp, n: n}
}

func (rw *RW) randomPick() int {
	// get a random number form [1,accSum[n-1]]
	num := rand.Intn(rw.accSum[rw.n-1]) + 1
	// use binary search to find the pos of num
	// for example with arr: [1,3] means 1/4 chances pick 0 and 3/4 pick 1
	left, right := 0, rw.n-1
	for left <= right {
		mid := left + (right-left)/2
		if rw.accSum[mid] == num {
			return mid
		} else if rw.accSum[mid] < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
