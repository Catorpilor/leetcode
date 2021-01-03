package rn

import (
	"math/rand"

	"github.com/catorpilor/leetcode/utils"
)

type Solution struct {
	storage []int
	n       int
}

func Constructor(head *utils.ListNode) *Solution {
	cur := head
	var st []int
	for cur != nil {
		st = append(st, cur.Val)
		cur = cur.Next
	}

	return &Solution{storage: st, n: len(st)}
}

func (this *Solution) GetRandom() int {
	pos := rand.Intn(this.n)
	return this.storage[pos]
}

type Solution2 struct {
	head *utils.ListNode
}

func Constructor2(head *utils.ListNode) *Solution2 {
	return &Solution2{
		head: head
	}
}

// GetRandome2 use reservoir sampling with k = 1. time complexity O(N)
func (this *Solution2) GetRandom2() int {
	scope := 1
	var ans int
	cur := this.head
	for cur != nil {
		if rand.Float32() < (1.0 / float32(scope)) {
			ans = cur.Val
		}
		cur = cur.Next
		scope++
	}
	return ans
}
