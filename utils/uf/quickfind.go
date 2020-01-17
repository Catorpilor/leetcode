package uf

type QuickFind struct {
    store []int
}

func NewQuickFind(n int) *QuickFind {
    st := make([]int, n)
    for i := range st {
        st[i] = i
    }
    return &QuickFind{store: st}
}

// Find time complexity O(1)
func (qf *QuickFind) Find(i, j int) bool {
    return qf.store[i] == qf.store[j]
}

// Union time complexity O(N)
func (qf *QuickFind) Union(i, j int) {
    qi := qf.store[i]
    for i := range qf.store {
        if qf.store[i] == qi {
            qf.store[i] = qf.store[j]
        }
    }
}
