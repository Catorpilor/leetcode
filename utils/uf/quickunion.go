package uf

type QuickUnion struct {
    store []int
}

func NewQuickUnion(n int) *QuickUnion {
    st := make([]int, n)
    for i := range st {
        st[i] = i
    }
    return &QuickUnion{store: st}
}

func (qu *QuickUnion) Find(i, j int) bool {
    return qu.root(i) == qu.root(j)
}

func (qu *QuickUnion) Union(i, j int) {
    ri, rj := qu.root(i), qu.root(j)
    qu.store[ri] = rj
}

func (qu *QuickUnion) root(i int) int {
    for qu.store[i] != i {
        i = qu.store[i]
    }
    return i
}
