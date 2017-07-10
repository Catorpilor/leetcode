package minstack

type MinStack2 struct {
	s   []int
	min int
}

func Constructor2() MinStack2 {
	return MinStack2{}
}

func (ms *MinStack2) Push(x int) {
	// just append the gap between x and min
	if len(ms.s) == 0 {
		ms.s = append(ms.s, 0)
		ms.min = x
	} else {
		ms.s = append(ms.s, x-ms.min)
		if x < ms.min {
			ms.min = x
		}
	}
}

func (ms *MinStack2) Pop() {
	if len(ms.s) == 0 {
		return
	}
	l := len(ms.s)
	t := ms.s[l-1]
	if t < 0 {
		ms.min = ms.min - t
	}
	ms.s = ms.s[:l-1]
}

func (ms *MinStack2) Top() int {
	l := len(ms.s)
	if l == 0 {
		panic("no values inside")
	}
	t := ms.s[l-1]
	if t > 0 {
		return ms.min + t
	} else {
		return ms.min
	}
}

func (ms *MinStack2) getMin() int {
	return ms.min
}
