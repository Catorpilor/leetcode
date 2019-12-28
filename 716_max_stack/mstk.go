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

type MaxStack2 struct {
	list, maxList []int
}

func NewMaxStack2() *MaxStack2 {
	return &MaxStack2{
		list:    make([]int, 0),
		maxList: make([]int, 0),
	}
}

func (this *MaxStack2) Push(x int) {
	this.list = append(this.list, x)
	if len(this.maxList) == 0 || x >= this.PeekMax() {
		this.maxList = append(this.maxList, x)
	}
}

func (this *MaxStack2) PeekMax() int {
	return this.maxList[len(this.maxList)-1]
}

func (this *MaxStack2) Top() int {
	return this.list[len(this.list)-1]
}

func (this *MaxStack2) Pop() int {
	v := this.Top()
	if v == this.PeekMax() {
		this.maxList = this.maxList[:len(this.maxList)-1]
	}
	this.list = this.list[:len(this.list)-1]
	return v
}

func (this *MaxStack2) PopMax() int {
	v := this.PeekMax()
	tmp := make([]int, 0, len(this.list))
	for this.Top() != v {
		tmp = append(tmp, this.Pop())
	}
	this.Pop()
	for i := len(tmp) - 1; i >= 0; i-- {
		this.Push(tmp[i])
	}
	return v
}
