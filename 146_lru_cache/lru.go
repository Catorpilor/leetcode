package lru

type DBNode struct {
	Key, Val   int
	Prev, Next *DBNode
}

type LRUCache struct {
	cap, n               int
	dummyHead, dummyTail *DBNode
	set                  map[int]*DBNode
}

func Constructor(d int) *LRUCache {
	head, tail := &DBNode{}, &DBNode{}
	head.Next = tail
	tail.Prev = head
	return &LRUCache{
		cap:       d,
		dummyHead: head,
		dummyTail: tail,
		set:       make(map[int]*DBNode, d),
	}
}

func (cache *LRUCache) Get(key int) int {
	if v, exists := cache.set[key]; exists {
		cache.moveFront(v)
		return v.Key
	}
	return -1
}

func (cache *LRUCache) Put(key, val int) {
	if old, exists := cache.set[key]; !exists {
		oldNext := cache.dummyHead.Next
		node := &DBNode{Val: val, Key: key, Prev: cache.dummyHead, Next: oldNext}
		cache.dummyHead.Next = node
		oldNext.Prev = node
		if cache.n == cache.cap {
			ct := cache.dummyTail.Prev
			tk := ct.Key
			delete(cache.set, tk)
			ct.Prev.Next = ct.Next
			ct.Next.Prev = ct.Prev
			cache.n--
		}
		cache.n++
		cache.set[key] = node
	} else {
		old.Val = val
		cache.moveFront(old)
	}
	return
}

func (cache *LRUCache) moveFront(node *DBNode) {
	if node != cache.dummyHead.Next {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		node.Prev = cache.dummyHead
		node.Next = cache.dummyHead.Next
		cache.dummyHead.Next = node
	}
}
