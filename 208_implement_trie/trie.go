package trie

// Trie implemented using ternary search trie
type Trie struct {
	Root *TstNode
}

type TstNode struct {
	Value               int
	Cb                  byte
	Left, Middle, Right *TstNode
}

func Constructor() *Trie {
	return &Trie{}
}

func put(node *TstNode, key string, value, dep int) *TstNode {
	c := key[dep]
	if node == nil {
		node = &TstNode{Cb: c}
	}
	if c < node.Cb {
		node.Left = put(node.Left, key, value, dep)
	} else if c > node.Cb {
		node.Right = put(node.Right, key, value, dep)
	} else if dep < len(key)-1 {
		node.Middle = put(node.Middle, key, value, dep+1)
	} else {
		node.Value = value
	}
	return node
}

func (t *Trie) Insert(s string) {
	t.Root = put(t.Root, s, 1, 0)
}

func get(node *TstNode, key string, dep int) *TstNode {
	if node == nil || key == "" {
		return nil
	}
	c := key[dep]
	if c < node.Cb {
		return get(node.Left, key, dep)
	} else if c > node.Cb {
		return get(node.Right, key, dep)
	} else if dep < len(key)-1 {
		return get(node.Middle, key, dep+1)
	}
	return node
}

func (t *Trie) Search(s string) bool {
	node := get(t.Root, s, 0)
	if node == nil || node.Value == 0 {
		return false
	}
	return true
}

func prefixWithNode(node *TstNode) bool {
	if node == nil {
		return false
	}
	if node.Value != 0 {
		return true
	}
	return prefixWithNode(node.Left) || prefixWithNode(node.Middle) || prefixWithNode(node.Right)
}

func (t *Trie) StartWith(prefix string) bool {
	node := get(t.Root, prefix, 0)
	if node == nil {
		return false
	}
	if node.Value != 0 {
		return true
	}
	return prefixWithNode(node.Middle)
}
