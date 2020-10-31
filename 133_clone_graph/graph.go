package graph

// Node represent a graph node
type Node struct {
	Val       int
	Neighbors []*Node
}

func isEqual(a, b *node) bool {
	return false
}

func genNode(adj [][]int) *Node {
	return nil
}

func cloneGraph(a *Node) *Node {
	set := make(map[int]*Node)
	return useDFS(set, a)
}

// useDFS time complexity O(N), space complexity O(N)
func useDFS(set map[int]*Node, node *Node) *Node {
	if node == nil {
		return nil
	}
	if zn, exists := set[node.Val]; exists {
		return zn
	}
	nn := &Node{Val: node.Val}
	for _, nb := range node.Neighbors {
		nn.Neighbors = append(nn.Neighbors, useDFS(set, nb))
	}
	set[node.Val] = nn
	return nn
}
