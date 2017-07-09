package minstack

type MinStack struct {
	s   []int // use slice s as stack
	min int   // current min value
}

func Constructor() MinStack {
	return MinStack{
		min: int(^uint(0) >> 1),
	}
}

func (m *MinStack) Push(x int) {
	// if x is less than current min
	// then push current min into the stack
	if x <= m.min {
		m.s = append(m.s, m.min)
		m.min = x
	}
	m.s = append(m.s, x)
}

func (m *MinStack) Pop() {
	l := len(m.s)
	if m.s[l-1] == m.min {
		// if top is m.min
		// pop twice
		m.min = m.s[l-2]
		m.s = m.s[:l-2]
	} else {
		m.s = m.s[:l-1]
	}
}

func (m *MinStack) Top() int {
	return m.s[len(m.s)-1]
}

func (m *MinStack) getMin() int {
	return m.min
}
