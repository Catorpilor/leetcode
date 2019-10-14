package swap

import "github.com/catorpilor/leetcode/utils"

func swapPairs(head *utils.ListNode) *utils.ListNode {
    if head == nil {
        return nil
    }
    p1, p2 := head, head.Next
    var temp int
    for p2 != nil {
        // swap adjacent
        temp = p1.Val
        p1.Val = p2.Val
        p2.Val = temp

        // update pointer
        p1 = p2.Next
        if p1 != nil && p1.Next != nil {
            p2 = p1.Next
        } else {
            p2 = nil
        }

    }

    return head
}
