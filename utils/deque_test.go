package utils

import "testing"

func newDequeWithtest(t *testing.T) *Deque {
    dep := NewDeque()
    t.Cleanup(func() {
        dep = nil
    })
    return dep
}

func TestIsEmpty(t *testing.T) {
    deq := newDequeWithtest(t)
    if deq.IsEmpty() != true {
        t.Fatal("with empty deque expected true but got false")
    }
    deq.PushBack(1)
    if deq.IsEmpty() != false {
        t.Fatal("after PushBack  deque.IsEmpty expect false but got true")
    }
}

func TestPushback(t *testing.T) {
    deq := newDequeWithtest(t)
    deq.PushBack(1)
    deq.PushBack(2)
    got := deq.Back().(int)
    if got != 2 {
        t.Fatalf("after pushback(2) expected 2 but got %d", got)
    }
    deq.PushBack(3)
    got = deq.Back().(int)
    if got != 3 {
        t.Fatalf("after pushback(3) expected 3 but got %d", got)
    }
}

func TestPushfront(t *testing.T) {
    deq := newDequeWithtest(t)
    deq.PushFront(1)
    deq.PushFront(2)
    got := deq.Front().(int)
    if got != 2 {
        t.Fatalf("after pushfront(2) expected 2 but got %d", got)
    }
    deq.PushFront(3)
    got = deq.Front().(int)
    if got != 3 {
        t.Fatalf("after pushfront(3) expected 3 but got %d", got)
    }
}

func TestPopback(t *testing.T) {
    deq := newDequeWithtest(t)
    deq.PushFront(1)
    deq.PushFront(2)
    deq.PushBack(3)
    if deq.n != 3 {
        t.Fatalf("after 3 push should be 3 elements but got %d", deq.n)
    }
    got := deq.PopBack().(int)
    if got != 3 {
        t.Fatalf("deq.PopBack expected 3 but got %d", got)
    }
    if deq.n != 2 {
        t.Fatalf("after 1 pop should be 2 elements but got %d", deq.n)
    }
    got = deq.Back().(int)
    if got != 1 {
        t.Fatalf("after popback expected back to be 2 but got %d", got)
    }
}

func TestPopfront(t *testing.T) {
    deq := newDequeWithtest(t)
    deq.PushFront(1)
    deq.PushFront(2)
    deq.PushBack(3)
    got := deq.PopFront().(int)
    if got != 2 {
        t.Fatalf("deq.PopFront expected 2 but got %d", got)
    }
    got = deq.Front().(int)
    if got != 1 {
        t.Fatalf("after popfront expected front to be 1 but got %d", got)
    }
}
