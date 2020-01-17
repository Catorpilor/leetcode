package uf

// WeightedQuickUnion M union-find ops on a N objects
// worst time complexity is o(N+MlgN)
type WeightedQuickUnion struct {
    store []int
    sz    []int
}

func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
    st := make([]int, n)
    for i := range st {
        st[i] = i
    }
    return &WeightedQuickUnion{store: st, sz: make([]int, n)}
}

// Find time complexity proportional to depth of i/j, depth is at most lgN
// worst case O(lgN)
func (qu *WeightedQuickUnion) Find(i, j int) bool {
    return qu.root(i) == qu.root(j)
}

// Union just like Find, time complexity Worst Case O(lgN)
func (qu *WeightedQuickUnion) Union(i, j int) {
    ri, rj := qu.root(i), qu.root(j)
    if qu.sz[ri] < qu.sz[rj] {
        qu.store[ri] = rj
        qu.sz[rj] += qu.sz[ri]
    } else {
        qu.store[rj] = ri
        qu.sz[ri] += qu.sz[rj]
    }
}

func (qu *WeightedQuickUnion) root(i int) int {
    for qu.store[i] != i {
        i = qu.store[i]
    }
    return i
}
