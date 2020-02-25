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

// useDFS time complexity O(N), space complexity O(N)
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

// useStack time complexity O(N) space complexity O(N)
func useStack(head *Node) *Node {
	if head == nil {
		return head
	}
	dummy := &Node{Next: head}
	stack := make([]*Node, 0, 1000)
	prev, cur := dummy, dummy
	stack = append(stack, head)
	n := len(stack)
	for n > 0 {
		cur = stack[n-1]
		stack = stack[:n-1]
		prev.Next = cur
		cur.Prev = prev
		if cur.Next != nil {
			// right sub tree is not nil
			stack = append(stack, cur.Next)
		}
		if cur.Child != nil {
			// left sub tree
			stack = append(stack, cur.Child)
			cur.Child = nil
		}
		prev = cur
		n = len(stack)
	}
	return dummy.Next
}
