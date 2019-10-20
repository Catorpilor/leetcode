package ugly

import (
	"container/heap"

	"github.com/catorpilor/leetcode/utils"
)

func NthUglyNumber(n int) int {
	// time exceeds
	if n < 1 {
		return -1
	}
	dp := []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12}
	if n > 10 {
		t := make([]int, n-10)
		dp = append(dp, t...)
	}
	for i, j := 10, 13; i < n; j++ {
		if isUgly(j) {
			dp[i] = j
			i += 1
		}
	}

	return dp[n-1]

}

func isUgly(n int) bool {
	for i := 2; i < 6; i++ {
		for n%i == 0 {
			n /= i
		}
	}
	return n == 1
}

func NthUglyNumber2(n int) int {
	if n < 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	var p2, p3, p5 int
	for i := 1; i < n; i++ {
		dp[i] = utils.Min(dp[p2]*2, utils.Min(dp[p3]*3, dp[p5]*5))
		if dp[i] == dp[p2]*2 {
			p2 += 1
		}
		if dp[i] == dp[p3]*3 {
			p3 += 1
		}
		if dp[i] == dp[p5]*5 {
			p5 += 1
		}
	}
	return dp[n-1]
}

func NthUglyNumber3(n int) int {
	if n < 1 {
		return 0
	}
	return useHeap(n)
}

func useHeap(n int) int {
	res := make([]int, 1690)
	pq := &utils.Int64PriorityQueue{}
	seen := make(map[int64]bool)
	heap.Push(pq, int64(1))
	seen[int64(1)] = true
	primes := []int64{2, 3, 5}
	var cur, gen int64
	for i := 0; i < 1690; i++ {
		cur = heap.Pop(pq).(int64)
		res[i] = int(cur)
		for _, p := range primes {
			gen = p * cur
			if !seen[gen] {
				seen[gen] = true
				heap.Push(pq, gen)
			}
		}
	}
	return res[n-1]
}
