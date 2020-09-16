package xor

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func maxXor(nums []int) int {
	return useTrie(nums)
}

// useBruteForce time complexity O(N^2), space complexity O(1)
func useBruteForce(nums []int) int {
	var ans int
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			ans = utils.Max(ans, nums[i]^nums[j])
		}
	}
	return ans
}

// useTrie time complexity O(N), space compplexity O(N)
func useTrie(nums []int) int {
	type trie struct {
		children []*trie
	}
	root := &trie{children: make([]*trie, 2)}
	for _, num := range nums {
		cur := root
		for i := 30; i >= 0; i-- {
			curBit := (num >> i) & 1
			if cur.children[curBit] == nil {
				cur.children[curBit] = &trie{children: make([]*trie, 2)}
			}
			cur = cur.children[curBit]
		}
	}
	ans := math.MinInt32
	for _, num := range nums {
		cur := root
		curMax := 0
		for i := 30; i >= 0; i-- {
			curBit := (num >> i) & 1
			if cur.children[curBit^1] != nil {
				curMax += (1 << i)
				cur = cur.children[curBit^1]
			} else {
				cur = cur.children[curBit]
			}
		}
		ans = utils.Max(ans, curMax)
	}
	return ans
}
