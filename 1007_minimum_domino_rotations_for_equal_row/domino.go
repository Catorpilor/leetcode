package domino

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func minRotations(a, b []int) int {
	return useBruteForce(a, b)
}

// useBruteForce time complexity O(N), space complexity O(1)
func useBruteForce(a, b []int) int {
	n := len(a)
	var ans int
	bucket1, bucket2 := make([]int, 7), make([]int, 7)
	for i := range a {
		bucket1[a[i]]++
		bucket2[b[i]]++
	}
	p1, p2 := 1, 1
	for i := range bucket1 {
		if bucket1[i] == n || bucket2[i] == n {
			return ans
		}
		if bucket1[i] > bucket1[p1] {
			p1 = i
		}
		if bucket2[i] > bucket2[p2] {
			p2 = i
		}
	}
	var r1, r2 int
	for i := range a {
		if a[i] != p1 {
			if b[i] == p1 {
				r1++
			} else {
				r1 = math.MaxInt32
				break
			}
		}
	}
	for i := range b {
		if b[i] != p2 {
			if a[i] == p2 {
				r2++
			} else {
				r2 = math.MaxInt32
				break
			}
		}
	}
	ans = utils.Min(r1, r2)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// useDiff time complexity O(N), space complexity O(1)
func useDiff(a, b []int) int {
	n := len(a)
	bkt1, bkt2, same := make([]int, 7), make([]int, 7), make([]int, 7)
	for i := range a {
		bkt1[a[i]]++
		bkt2[b[i]]++
		if a[i] == b[i] {
			same[a[i]]++
		}
	}
	for i := 1; i < 7; i++ {
		if bkt1[i]+bkt2[i]-same[i] == n {
			return n - utils.Max(bkt1[i], bkt2[i])
		}
	}
	return -1
}
