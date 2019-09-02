package subarr

import "fmt"

func subarraysWithKDistinct(A []int, K int) int {
	n := len(A)
	if n < 1 || K < 1 || K > n {
		return 0
	}
	// exeactly K means
	// atMost(A, K) - atMost(A, K-1)
	return atMost(A, K) - atMost(A, K-1)
}

func atMost(A []int, K int) int {
	var beg, end, res, disC int
	n := len(A)
	cache := make(map[int]int, n)
	for end < n {
		cache[A[end]]++
		if cache[A[end]] == 1 {
			disC++
		}
		end++
		for disC > K {
			cache[A[beg]]--
			if cache[A[beg]] == 0 {
				disC--
			}
			beg++
		}
		// from (beg,end) there are end-beg subarrays that have at most K distinct numbers
		// for example K = 2, A=[1,2,1,3]
		// atMost(A,2) the avaliable subarrays are
		// [1],[2],[1],[3],[1,2],[1,2,1],[2,1],[1,3]
		// atMost(A,1) the avaliable subarrays are
		//[1],[2],[1],[3]
		// fmt.Printf("beg: %d, end: %d ,disC: %d, res: %d\n", beg, end, disC, res)
		res += end - beg
		fmt.Printf("beg: %d, end: %d ,disC: %d, res: %d\n", beg, end, disC, res)
	}
	return res
}
