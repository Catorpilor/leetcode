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
