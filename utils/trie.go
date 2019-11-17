package utils

// TSTNode represents the ternary search trie node
type TSTNode struct {
	value               interface{}
	cb                  byte
	left, middle, right *TSTNode
}

type TernarySearchTrie struct {
	root *TSTNode
}

func NewTST() *TernarySearchTrie {
	return &TernarySearchTrie{
		root: &TSTNode{},
	}
}

func (t *TernarySearchTrie) Put(key string, val interface{}) {
	t.root = put(t.root, key, val, 0)
}

func put(node *TSTNode, key string, val interface{}, d int) *TSTNode {
	c := key[d]
	if node == nil {
		node = &TSTNode{
			cb: c,
		}
	}
	if c < node.cb {
		//fmt.Printf("node.cb: %c, and c: %c,go to left\n", node.cb, c)
		node.left = put(node.left, key, val, d)
	} else if c > node.cb {
		//fmt.Printf("node.cb: %c, and c: %c,go to right\n", node.cb, c)
		node.right = put(node.right, key, val, d)
	} else if d < len(key)-1 {
		//fmt.Printf("node.cb: %c, and c: %c,go to middle\n", node.cb, c)
		node.middle = put(node.middle, key, val, d+1)
	} else {
		//fmt.Printf("node.c: %c, d=%d, len(key):%d\n", node.cb, d, len(key))
		node.value = val
	}
	return node
}

func (t *TernarySearchTrie) LongestPrefixOf(query string) string {
	length := search(t.root, query, 0, 0)
	return query[:length]
}

func search(node *TSTNode, query string, d, length int) int {
	if node == nil {
		return length
	}
	if node.value != nil {
		length = d + 1 // d+1 means the actually length
	}
	if d == len(query) {
		return length
	}
	c := query[d]
	if c < node.cb {
		// fmt.Printf("node.cb: %c, and c: %c,go to left\n", node.cb, c)
		return search(node.left, query, d, length)
	} else if c > node.cb {
		// fmt.Printf("node.cb: %c, and c: %c,go to right\n", node.cb, c)
		return search(node.right, query, d, length)
	} else if d < len(query) {
		// fmt.Printf("node.cb: %c, and c: %c,go to middle with d:%d, and query:%d\n", node.cb, c, d, len(query))
		return search(node.middle, query, d+1, length)
	}
	return length
}

func (t *TernarySearchTrie) Contains(key string) bool {
	return t.Get(key) != nil
}

func (t *TernarySearchTrie) Get(key string) interface{} {
	x := get(t.root, key, 0)
	if x == nil {
		return nil
	}
	return x.value
}

func get(node *TSTNode, key string, d int) *TSTNode {
	if node == nil {
		return nil
	}
	c := key[d]
	if c < node.cb {
		return get(node.left, key, d)
	} else if c > node.cb {
		return get(node.right, key, d)
	} else if d < len(key)-1 {
		return get(node.middle, key, d+1)
	} else {
		return node
	}
}
