package flat

// Node represents a multi level doubly linked list node
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// func constructFromSlice([]int) *Node {
// 	return nil
// }

func flatten(head *Node) *Node {
	return useDFS(head)
}

func useDFS(head *Node) *Node {
	if head == nil {
		return head
	}
	dummy := &Node{Next: head}
	dfs(dummy, head)
	dummy.Next.Prev = nil
	return dummy.Next
}

// dfs preOrder traversal
func dfs(prev, cur *Node) *Node {
	if cur == nil {
		return prev
	}
	// link prev and cur
	prev.Next = cur
	cur.Prev = prev

	next := cur.Next
	// left sub tree
	newHead := dfs(cur, cur.Child)
	cur.Child = nil
	// right sub tree
	return dfs(newHead, next)
}
