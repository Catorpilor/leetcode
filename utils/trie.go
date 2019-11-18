package utils

// TSTNode represents the ternary search trie node
type TSTNode struct {
	Value               interface{}
	Cb                  byte
	Left, Middle, Right *TSTNode
}

type TernarySearchTrie struct {
	Root *TSTNode
}

func NewTST() *TernarySearchTrie {
	return &TernarySearchTrie{
		//	Root: &TSTNode{}, not a R-way trie, root is not empty
	}
}

func (t *TernarySearchTrie) Put(key string, val interface{}) {
	t.Root = put(t.Root, key, val, 0)
}

func put(node *TSTNode, key string, val interface{}, d int) *TSTNode {
	c := key[d]
	if node == nil {
		node = &TSTNode{
			Cb: c,
		}
	}
	if c < node.Cb {
		//fmt.Printf("node.Cb: %c, and c: %c,go to left\n", node.Cb, c)
		node.Left = put(node.Left, key, val, d)
	} else if c > node.Cb {
		//fmt.Printf("node.Cb: %c, and c: %c,go to right\n", node.Cb, c)
		node.Right = put(node.Right, key, val, d)
	} else if d < len(key)-1 {
		//fmt.Printf("node.Cb: %c, and c: %c,go to Middle\n", node.Cb, c)
		node.Middle = put(node.Middle, key, val, d+1)
	} else {
		//fmt.Printf("node.c: %c, d=%d, len(key):%d\n", node.Cb, d, len(key))
		node.Value = val
	}
	return node
}

func (t *TernarySearchTrie) LongestPrefixOf(query string) string {
	length := search(t.Root, query, 0, 0)
	return query[:length]
}

func search(node *TSTNode, query string, d, length int) int {
	if node == nil {
		return length
	}
	if node.Value != nil {
		length = d + 1 // d+1 means the actually length
	}
	if d == len(query) {
		return length
	}
	c := query[d]
	if c < node.Cb {
		// fmt.Printf("node.Cb: %c, and c: %c,go to Left\n", node.Cb, c)
		return search(node.Left, query, d, length)
	} else if c > node.Cb {
		// fmt.Printf("node.Cb: %c, and c: %c,go to Right\n", node.Cb, c)
		return search(node.Right, query, d, length)
	} else if d < len(query) {
		// fmt.Printf("node.Cb: %c, and c: %c,go to Middle with d:%d, and query:%d\n", node.Cb, c, d, len(query))
		return search(node.Middle, query, d+1, length)
	}
	return length
}

func (t *TernarySearchTrie) Contains(key string) bool {
	return t.Get(key) != nil
}

func (t *TernarySearchTrie) Get(key string) interface{} {
	x := get(t.Root, key, 0)
	if x == nil {
		return nil
	}
	return x.Value
}

func get(node *TSTNode, key string, d int) *TSTNode {
	if node == nil {
		return nil
	}
	c := key[d]
	if c < node.Cb {
		return get(node.Left, key, d)
	} else if c > node.Cb {
		return get(node.Right, key, d)
	} else if d < len(key)-1 {
		return get(node.Middle, key, d+1)
	} else {
		return node
	}
}
