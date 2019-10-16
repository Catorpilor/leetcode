package mksl

import (
    "math"
    "sort"

    "github.com/catorpilor/leetcode/utils"
)

func mergeKLists(lists []*utils.ListNode) *utils.ListNode {
    // return bruteForce(lists)
    return compare(lists)
}

func bruteForce(lists []*utils.ListNode) *utils.ListNode {
    nodes := make([]int, 0, 10000)
    for _, l := range lists {
        for l != nil {
            nodes = append(nodes, l.Val)
            l = l.Next
        }
    }
    if len(nodes) > 1 {
        sort.Slice(nodes, func(i, j int) bool {
            return nodes[i] <= nodes[j]
        })
    }
    return utils.ConstructFromSlice(nodes)
}

func compare(lists []*utils.ListNode) *utils.ListNode {
    k := len(lists)
    pointers := make([]*utils.ListNode, k)
    for i := range pointers {
        pointers[i] = lists[i]
    }
    sortedNodes := make([]int, 0, 1000000)
    for {
        sm, idx := math.MaxInt32, -1
        for i := range pointers {
            if pointers[i] != nil && pointers[i].Val < sm {
                sm = pointers[i].Val
                idx = i
            }
        }
        if idx == -1 {
            // we are done
            break
        }
        sortedNodes = append(sortedNodes, pointers[idx].Val)
        // update pointers[idx]
        pointers[idx] = pointers[idx].Next
    }
    return utils.ConstructFromSlice(sortedNodes)
}
