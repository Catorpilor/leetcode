package character

type trieNode struct {
	isWord bool
	next   []*trieNode
}

// StreamChecker
type StreamChecker struct {
	root *trieNode
	q    []byte
}

// COnstructor create a StreamChecker based on words
func Constructor(words []string) *StreamChecker {
	root := createTrie(words)
	return &StreamChecker{
		root: root,
		q:    make([]byte, 0, 40000),
	}
}

func createTrie(words []string) *trieNode {
	n := len(words)
	if n < 1 {
		return nil
	}
	root := &trieNode{next: make([]*trieNode, 26)}
	for _, w := range words {
		nw := len(w)
		node := root
		for i := nw - 1; i >= 0; i-- {
			c := w[i]
			if node.next[c-'a'] == nil {
				node.next[c-'a'] = &trieNode{next: make([]*trieNode, 26)}
			}
			node = node.next[c-'a']
		}
		node.isWord = true
	}
	return root
}

// Query query a letter
func (sc *StreamChecker) Query(letter byte) bool {
	sc.q = append(sc.q, letter)
	n := len(sc.q)
	node := sc.root
	for i := n - 1; i >= 0 && node != nil; i-- {
		c := sc.q[i]
		node = node.next[c-'a']
		if node != nil && node.isWord {
			return true
		}
	}
	return false
}
