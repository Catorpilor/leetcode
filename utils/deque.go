package utils

// DNode represents the double linked node
type DNode struct {
    Val        interface{}
    Prev, Next *DNode
}

// Deque is the double ended queue
type Deque struct {
    head, tail *DNode
    n          int
}

func NewDeque() *Deque {
    head, tail := &DNode{}, &DNode{}
    head.Next = tail
    tail.Prev = head

    return &Deque{
        head: head,
        tail: tail,
    }
}

// IsEmpty check deque is empty or not
func (deq *Deque) IsEmpty() bool {
    return deq.n == 0
}

// PushFront push val at front
func (deq *Deque) PushFront(val interface{}) {
    node := &DNode{Val: val, Prev: deq.head, Next: deq.head.Next}
    deq.head.Next.Prev = node
    deq.head.Next = node
    deq.n++
}

// PushBack push val at back
func (deq *Deque) PushBack(val interface{}) {
    node := &DNode{Val: val, Next: deq.tail, Prev: deq.tail.Prev}
    deq.tail.Prev.Next = node
    deq.tail.Prev = node
    deq.n++
}

func (deq *Deque) PopBack() interface{} {
    if deq.n == 0 {
        return nil
    }
    node := deq.tail.Prev
    node.Prev.Next = deq.tail
    deq.tail.Prev = node.Prev
    deq.n--
    return node.Val
}

func (deq *Deque) PopFront() interface{} {
    if deq.n == 0 {
        return nil
    }
    node := deq.head.Next
    node.Next.Prev = deq.head
    deq.head.Next = node.Next
    deq.n--
    return node.Val
}

func (deq *Deque) Back() interface{} {
    if deq.n == 0 {
        return nil
    }
    return deq.tail.Prev.Val
}

func (deq *Deque) Front() interface{} {
    if deq.n == 0 {
        return nil
    }
    return deq.head.Next.Val
}
