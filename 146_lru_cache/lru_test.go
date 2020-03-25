package lru

import "testing"

var obj *LRUCache

func setup(cap int) func() {
	if obj == nil {
		obj = Constructor(cap)
	}
	return func() { obj = nil }
}

func TestEvict(t *testing.T) {
	f := setup(2)
	defer f()
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3) // 1,1 should evicted
	if obj.Get(1) != -1 {
		t.Fatal("with cap(2) insert the 3rd element should evict the oldest one")
	}
	t.Log("pass")
}

func TestWithBunchCommands(t *testing.T) {
	f := setup(2)
	defer f()
	obj.Put(1, 1)
	obj.Put(2, 2)
	if obj.Get(1) != 1 {
		t.Fatalf("after put (1,1), get should return 1 but got:%d", obj.Get(1))
	}
	obj.Put(3, 3)
	if obj.Get(2) != -1 {
		t.Fatalf("after put (3,3), get should return -1 but got:%d", obj.Get(2))
	}
	obj.Put(4, 4)
	if obj.Get(1) != -1 {
		t.Fatalf("after put (4,4), get should return -1 but got:%d", obj.Get(1))
	}
	// 2,2 is the oldest data now, should be evicted
	if obj.Get(3) != 3 {
		t.Fatalf("call get(3), get should return 3 but got:%d", obj.Get(3))
	}
	if obj.Get(4) != 4 {
		t.Fatalf("call get(4), get should return 4 but got:%d", obj.Get(4))
	}
	t.Log("pass")
}
