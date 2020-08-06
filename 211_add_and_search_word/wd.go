package wd

import "fmt"

type trieNode struct {
	Value               int
	Cb                  byte
	Left, Middle, Right *trieNode
}

func put(node *trieNode, key string, value, d int) *trieNode {
	c := key[d]
	if node == nil {
		node = &trieNode{Cb: c}
	}
	if c < node.Cb {
		// go to left node
		fmt.Printf("node.Cb: %c, and c:%c, go to left\n", node.Cb, c)
		node.Left = put(node.Left, key, value, d)
	} else if c > node.Cb {
		// go to right node
		fmt.Printf("node.Cb:%c, c: %c go to right\n", node.Cb, c)
		node.Right = put(node.Right, key, value, d)
	} else if d < len(key)-1 {
		fmt.Printf("node.Cb:%c c: %c goto middle\n", node.Cb, c)
		node.Middle = put(node.Middle, key, value, d+1)
	} else {
		// reach the end of key
		fmt.Printf("update node's value:%v, c:%c, d:%d, len(key)=%d\n", value, c, d, len(key))
		node.Value = value
	}
	return node
}

func get(node *trieNode, key string, d int) bool {
	if key == "" || node == nil {
		return false
	}
	if d == len(key) {
		fmt.Printf("reach the end, with node.Val:%d\n", node.Value)
		return node.Value == len(key)
	}
	c := key[d]
	fmt.Printf("node.Cb:%c, c:%c, d:%d\n", node.Cb, c, d)
	if c != '.' {
		if c < node.Cb {
			return get(node.Left, key, d)
		} else if c > node.Cb {
			return get(node.Right, key, d)
		} else if d < len(key)-1 {
			return get(node.Middle, key, d+1)
		} else {
			return node.Value == len(key)
		}
	} else {
		// this is the bottle neck.
		// maybe.. just maybe tst is not the proper choice for this fuzzy search.
		if d < len(key)-1 {
			return get(node.Left, key, d+1) || get(node.Right, key, d+1) || get(node.Middle, key, d+1) || get(node.Left, key, d) || get(node.Right, key, d) || get(node.Middle, key, d)
		}
		return node.Value == len(key)
	}
	return false
}

type WordDict struct {
	root *trieNode
}

func Constructor() *WordDict {
	return &WordDict{}
}

func (this *WordDict) AddWord(word string) {
	this.root = put(this.root, word, len(word), 0)
}

func (this *WordDict) Search(word string) bool {
	return get(this.root, word, 0)
}
