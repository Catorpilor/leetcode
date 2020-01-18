package uf

// WQUPC M union-find ops on a N objects
// worst time complexity is o((N+M)*lgN)
type WQUPC struct {
    store []int
    sz    []int
}

func NewWQUPC(n int) *WQUPC {
    st := make([]int, n)
    sz := make([]int, n)
    for i := range st {
        st[i] = i
        sz[i] = 1
    }
    return &WQUPC{store: st, sz: sz}
}

// Find time complexity proportional to depth of i/j, depth is at most lgN
// worst case O(lgN)
func (qu *WQUPC) Find(i, j int) bool {
    return qu.Root(i) == qu.Root(j)
}

// Union just like Find, time complexity Worst Case O(lgN)
func (qu *WQUPC) Union(i, j int) {
    ri, rj := qu.Root(i), qu.Root(j)
    if ri == rj {
        return
    }
    if qu.sz[ri] < qu.sz[rj] {
        qu.store[ri] = rj
        qu.sz[rj] += qu.sz[ri]
    } else {
        qu.store[rj] = ri
        qu.sz[ri] += qu.sz[rj]
    }
}

func (qu *WQUPC) Root(i int) int {
    for qu.store[i] != i {
        // make every other nodes in path points to its grandparent
        qu.store[i] = qu.store[qu.store[i]]
        i = qu.store[i]
    }
    return i
}
