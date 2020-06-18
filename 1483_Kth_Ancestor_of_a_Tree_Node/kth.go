package kth

type TreeAncestor struct {
	// for now just use brute force
	n      int
	parent []int
}

func Construct(n int, parent []int) *TreeAncestor {
	return &TreeAncestor{
		n:      n,
		parent: parent,
	}
}

func (this *TreeAncestor) GetKthAncestor(node, k int) int {
	tmp := node
	for i := 1; i <= k; i++ {
		if this.parent[tmp] != -1 {
			tmp = this.parent[tmp]
		} else {
			return -1
		}
	}
	return tmp
}
