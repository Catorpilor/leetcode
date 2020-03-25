package lru

import (
	"fmt"
	"strings"
)

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
		// fmt.Printf("get key: %d, found\n", key)
		cache.moveFront(v)
		// fmt.Println(cache.String())
		return v.Key
	}
	// fmt.Printf("get key: %d, not found\n", key)
	// fmt.Println(cache.String())
	return -1
}

func (cache *LRUCache) Put(key, val int) {
	if old, exists := cache.set[key]; !exists {
		if cache.n == cache.cap {
			ct := cache.dummyTail.Prev
			// fmt.Printf("befor delete: %s, prepare to delete node: %+v\n", cache.String(), ct)
			tk := ct.Key
			delete(cache.set, tk)
			// fmt.Printf("deleting node: %p, prev(%p):%+v, next(%p):%+v\n", ct, ct.Prev, ct.Prev, ct.Next, ct.Next)
			ct.Prev.Next = ct.Next
			ct.Next.Prev = ct.Prev
			cache.n--
			// fmt.Println(cache.String())
		}
		oldNext := cache.dummyHead.Next
		node := &DBNode{Val: val, Key: key, Prev: cache.dummyHead, Next: oldNext}
		cache.dummyHead.Next = node
		oldNext.Prev = node
		cache.n++
		cache.set[key] = node
		// fmt.Printf("insert key:%d, not found\n", key)
		// fmt.Println(cache.String())
	} else {
		old.Val = val
		cache.moveFront(old)
		// fmt.Printf("insert key:%d, found\n", key)
		// fmt.Println(cache.String())
	}
	// fmt.Println(cache.String())
	return
}

func (cache *LRUCache) moveFront(node *DBNode) {
	// fmt.Printf("before move node(%p):%+v, cache:%s", node, node, cache.String())
	if node != cache.dummyHead.Next {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		node.Prev = cache.dummyHead
		node.Next = cache.dummyHead.Next
		cache.dummyHead.Next.Prev = node
		cache.dummyHead.Next = node
	}
	// fmt.Printf("after move node(%p):%+v, cache:%s", node, node, cache.String())
}

func (cache *LRUCache) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "cap:%d, n:%d, set:%v, list:0=", cache.cap, cache.n, cache.set)
	node := cache.dummyHead.Next
	for node != cache.dummyTail {
		fmt.Fprintf(&sb, "%d=", node.Val)
		node = node.Next
	}
	sb.WriteString("0")
	return sb.String()
}
