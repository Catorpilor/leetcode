package utils

type TrieNode struct {
	Value    interface{}
	Children []*TrieNode
}

const (
	defaultR = 26
)

type RTrie struct {
	Root *TrieNode
	r    int
}

func NewRTrie(r int) *RTrie {
	if r <= 0 {
		r = defaultR
	}
	return &RTrie{Root: &TrieNode{Children: make([]*TrieNode, r)}, r: r}
}

func (rt *RTrie) Put(w string) {
	p := rt.Root
	for _, c := range w {
		idx := c - 'a'
		if p.Children[idx] == nil {
			p.Children[idx] = &TrieNode{
				Children: make([]*TrieNode, rt.r),
			}
		}
		p = p.Children[idx]
	}
	p.Value = w
}

func (rt *RTrie) Contains(query string) bool {
	return rt.Get(query) != nil
}

func (rt *RTrie) Get(query string) interface{} {
	node := rget(rt.Root, query, 0)
	if node == nil {
		return nil
	}
	return node.Value
}

func rget(node *TrieNode, query string, d int) *TrieNode {
	if node == nil {
		return node
	}
	if d == len(query) {
		return node
	}
	idx := query[d] - 'a'
	return rget(node.Children[idx], query, d+1)
}
