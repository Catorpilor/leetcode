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
		node.left = put(node.left, key, val, d)
	} else if c > node.cb {
		node.right = put(node.right, key, val, d)
	} else if d < len(key)-1 {
		node.middle = put(node.middle, key, val, d+1)
	} else {
		node.value = val
	}
	return node
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
