package mstk

import "testing"

var mt *MaxStack

func setup() func() {
	if mt == nil {
		tmp := Constructor()
		mt = &tmp
	}
	// push
	res := []int{2, 1, 5, 3, 5, 4}
	for i := range res {
		mt.Push(res[i])
	}
	return func() { mt = nil }
}

func TestTop(t *testing.T) {
	f := setup()
	defer f()
	if mt.Top() != 4 {
		t.Fatalf("wanted 4 but got %d", mt.Top())
	}
	mt.Pop()
	if mt.Top() != 5 {
		t.Fatalf("wanted 5 but got %d", mt.Top())
	}
}

func TestPopMax(t *testing.T) {
	f := setup()
	defer f()
	if mt.PopMax() != 5 {
		t.Fatalf("wanted 5 but got %d", mt.PopMax())
	}
	if mt.PopMax() != 5 {
		t.Fatalf("wanted 5 but got %d", mt.PopMax())
	}
	if mt.PopMax() != 4 {
		t.Fatalf("wanted 4 but got %d", mt.PopMax())
	}

}

func TestPeekMax(t *testing.T) {
	setup()
	if mt.PeekMax() != 5 {
		t.Fatalf("wanted 5 but got %d", mt.PeekMax())
	}
	mt.PopMax()
	mt.PopMax()
	if mt.PeekMax() != 4 {
		t.Fatalf("wanted 4 but got %d", mt.PeekMax())
	}
}
