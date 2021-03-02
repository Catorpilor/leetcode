package stack

import "testing"

func constructWithTesting(t *testing.T) *FreqStack {
	fs := Construct()
	t.Cleanup(func() {
		fs = nil
	})
	return fs
}

func TestPush(t *testing.T) {
	fs := constructWithTesting(t)
	wanted := 7
	fs.Push(5)
	fs.Push(6)
	fs.Push(7)
	got := fs.Pop()
	if wanted != got {
		t.Fatalf("wanted %d but got %d", wanted, got)
	}
	t.Log("pass")
}

func TestPop(t *testing.T) {
	fs := constructWithTesting(t)
	fs.Push(5)
	fs.Push(7)
	fs.Push(5)
	fs.Push(7)
	fs.Push(4)
	fs.Push(5)
	pop1 := fs.Pop()
	if pop1 != 5 {
		t.Fatalf("wanted 5 but got %d", pop1)
	}
	pop1 = fs.Pop()
	if pop1 != 7 {
		t.Fatalf("wanted 7 but got %d", pop1)
	}
}
