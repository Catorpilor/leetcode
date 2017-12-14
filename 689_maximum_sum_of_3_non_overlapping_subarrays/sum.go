package sum

import "fmt"

func MaxSumOfThree(nums []int, k int) []int {
	n := len(nums)
	if k > n/3 || n/3 < 1 {
		return nil
	}
	// pre compute the accumlated sums
	sums := make([]int, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + nums[i]
	}
	fmt.Println(sums)

	// if the middle interval is [i,i+k-1],where k <= i <= n - 2k
	// then the left interval is [0,i-1]
	// the right interval is [i+k, n-1]

	// leftPos represents the starting index for the left interval in range [0,i-1]
	// rightPos represents the starting index for the right interval in range [i+k, n-1]
	leftPos, rightPos := make([]int, n), make([]int, n)
	for i := range rightPos {
		rightPos[i] = n - k
	}
	// find the left starting index
	// leftPos[i] means the left interval starting index when middle interval starting index is i
	for i, t := k, sums[k]-sums[0]; i <= n-2*k; i++ {
		if sums[i+1]-sums[i+1-k] > t {
			leftPos[i] = i + 1 - k
			t = sums[i+1] - sums[i+1-k]
		} else {
			leftPos[i] = leftPos[i-1]
		}
	}

	// find the right index
	for i, t := n-k-1, sums[n]-sums[n-k]; i >= 0; i-- {
		if sums[i+k]-sums[i] >= t {
			rightPos[i] = i
			t = sums[i+k] - sums[i]
		} else {
			rightPos[i] = rightPos[i+1]
		}
	}

	ret := make([]int, 3)
	var maxT int
	for i := k; i <= n-2*k; i++ {
		l, r := leftPos[i-1], rightPos[i+k]
		temp := (sums[i+k] - sums[i]) + (sums[l+k] - sums[l]) + sums[r+k] - sums[r]
		if temp > maxT {
			maxT = temp
			ret[0], ret[1], ret[2] = l, i, r
		}
	}
	return ret
}
