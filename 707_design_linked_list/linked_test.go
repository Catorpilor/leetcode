package linked

import "testing"

var list *MyLinkedList

func setup() func() {
	if list == nil {
		list = Constructor()
	}
	return func() { list = nil }
}

func TestAddAtIndex(t *testing.T) {
	f := setup()
	defer f()
	list.AddAtIndex(0, 1)
	list.AddAtIndex(0, 2)
	list.AddAtIndex(0, 3)
	if list.String() != "3->2->1" {
		t.Fatalf("after insert, list: %s should be 3->2->1", list)
	}
	list.AddAtIndex(3, 4)
	if list.String() != "3->2->1->4" {
		t.Fatalf("after insert, list: %s should be 3->2->1->4", list)
	}
}

func TestAddAtHead(t *testing.T) {
	f := setup()
	defer f()
	list.AddAtHead(1)
	list.AddAtHead(2)
	list.AddAtHead(3)
	if list.String() != "3->2->1" {
		t.Fatalf("after insert, list: %s should be 3->2->1", list)
	}
	list.AddAtHead(4)
	if list.String() != "4->3->2->1" {
		t.Fatalf("after insert, list: %s should be 4->3->2->1", list)
	}
}

func TestAddAtTail(t *testing.T) {
	f := setup()
	defer f()
	list.AddAtTail(1)
	list.AddAtTail(2)
	list.AddAtTail(3)
	if list.String() != "1->2->3" {
		t.Fatalf("after insert, list: %s should be 1->2->3", list)
	}
	list.AddAtTail(4)
	if list.String() != "1->2->3->4" {
		t.Fatalf("after insert, list: %s should be 1->2->3->4", list)
	}
}

func TestDeleteAtIndex(t *testing.T) {
	f := setup()
	defer f()
	list.AddAtTail(1)
	list.AddAtTail(2)
	list.AddAtTail(3)
	list.AddAtTail(6)
	list.AddAtTail(4)
	list.DeleteAtIndex(3)
	shoudBe := "1->2->3->4"
	if list.String() != shoudBe {
		t.Fatalf("after delete at index 3, list should be %s but got %s", shoudBe, list)
	}

}
