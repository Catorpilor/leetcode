package swap

import (
    "testing"

    "github.com/catorpilor/leetcode/utils"
)

func TestSwapPairs(t *testing.T) {
    st := []struct {
        name string
        head *utils.ListNode
        exp  *utils.ListNode
    }{
        {"empty list", nil, nil},
        {"single node", utils.ConstructFromSlice([]int{3}), utils.ConstructFromSlice([]int{3})},
        {"odd nodes", utils.ConstructFromSlice([]int{1, 2, 3}), utils.ConstructFromSlice([]int{2, 1, 3})},
        {"testcase1", utils.ConstructFromSlice([]int{1, 2, 3, 4}), utils.ConstructFromSlice([]int{2, 1, 4, 3})},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := swapPairs(tt.head)
            if !utils.IsEqualList(out, tt.exp) {
                t.Fatalf("with input list: %s wanted %s but got %s", tt.head, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
