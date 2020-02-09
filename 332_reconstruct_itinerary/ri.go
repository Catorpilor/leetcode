package ri

import (
	"container/heap"

	"github.com/catorpilor/leetcode/utils"
)

func ConstructIt(tickets [][]string) []string {
	lcs := make(map[string]*utils.StringHeap, len(tickets))
	for _, v := range tickets {
		if lcs[v[0]] == nil {
			lcs[v[0]] = utils.NewStringHeap()
		}
		heap.Push(lcs[v[0]], v[1])
	}
	st := utils.NewStack()
	st.Push("JFK")
	var ret []string
	for !st.IsEmpty() {
		for {
			if cp, exists := lcs[st.Top().(string)]; exists && cp.Len() > 0 {
				st.Push(heap.Pop(cp))
			} else {
				goto EXTLOOP
			}
		}
	EXTLOOP:
		ret = append([]string{st.Pop().(string)}, ret...)
	}
	return ret
}
