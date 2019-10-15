package mksl

import (
    "sort"

    "github.com/catorpilor/leetcode/utils"
)

func mergeKLists(lists []*utils.ListNode) *utils.ListNode {
    return bruteForce(lists)
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
