package mstk

import "testing"

var mt *MaxStack
var mt2 *MaxStack2

func setup() func() {
	if mt == nil {
		mt = Constructor()
	}
	// push
	res := []int{2, 1, 5, 3, 5, 4}
	for i := range res {
		mt.Push(res[i])
	}
	return func() { mt = nil }
}

func setup2() func() {
	if mt2 == nil {
		mt2 = NewMaxStack2()
	}
	// push
	res := []int{2, 1, 5, 3, 5, 4}
	for i := range res {
		mt2.Push(res[i])
	}
	return func() { mt2 = nil }
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
	if mt2.PeekMax() != 4 {
		t.Fatalf("wanted 4 but got %d", mt.PeekMax())
	}
}

func TestTop2(t *testing.T) {
	f := setup2()
	defer f()
	if mt2.Top() != 4 {
		t.Fatalf("wanted 4 but got %d", mt2.Top())
	}
	mt2.Pop()
	if mt2.Top() != 5 {
		t.Fatalf("wanted 5 but got %d", mt2.Top())
	}
}

func TestPopMax2(t *testing.T) {
	f := setup2()
	defer f()
	if mt2.PopMax() != 5 {
		t.Fatalf("wanted 5 but got %d", mt2.PopMax())
	}
	if mt2.PopMax() != 5 {
		t.Fatalf("wanted 5 but got %d", mt2.PopMax())
	}
	if mt2.PopMax() != 4 {
		t.Fatalf("wanted 4 but got %d", mt2.PopMax())
	}

}

func TestPeekMax2(t *testing.T) {
	f := setup2()
	defer f()
	if mt2.PeekMax() != 5 {
		t.Fatalf("wanted 5 but got %d", mt2.PeekMax())
	}
	mt2.PopMax()
	mt2.PopMax()
	if mt2.PeekMax() != 4 {
		t.Fatalf("wanted 4 but got %d", mt2.PeekMax())
	}
}
