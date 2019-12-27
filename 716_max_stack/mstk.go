package mstk

import (
	"github.com/catorpilor/leetcode/utils"
)

type MaxStack struct {
	list, maxList *utils.Stack
}

func Constructor() *MaxStack {
	return &MaxStack{
		list:    utils.NewStack(),
		maxList: utils.NewStack(),
	}
}

func (this *MaxStack) Push(x int) {
	curMax := x
	if !this.maxList.IsEmpty() {
		mtl := this.maxList.Top().(int)
		if mtl > x {
			curMax = mtl
		}
	}
	this.maxList.Push(curMax)
	this.list.Push(x)
}

func (this *MaxStack) Pop() int {
	_ = this.maxList.Pop()
	return this.list.Pop().(int)
}

func (this *MaxStack) Top() int {
	return this.list.Top().(int)
}

func (this *MaxStack) PeekMax() int {
	return this.maxList.Top().(int)
}

func (this *MaxStack) PopMax() int {
	// fmt.Printf("list: %s, max: %s\n", this.list, this.maxList)
	curMax := this.PeekMax()
	tmp := utils.NewStack()
	for this.Top() != curMax {
		tmp.Push(this.Pop())
	}
	this.Pop()
	for !tmp.IsEmpty() {
		this.Push(tmp.Pop().(int))
	}
	return curMax
}
