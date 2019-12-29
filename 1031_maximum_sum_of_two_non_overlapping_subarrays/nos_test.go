package nos

import "testing"

func TestMaxSum(t *testing.T) {
    st := []struct {
        name string
        A    []int
        l, m int
        exp  int
    }{
        {"empty array", []int{}, 1, 1, 0},
        {"testcase1", []int{0, 6, 5, 2, 2, 5, 1, 9, 4}, 1, 2, 20},
        {"testcase2", []int{3, 8, 1, 3, 2, 1, 8, 9, 0}, 3, 2, 29},
        {"testcase3", []int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}, 4, 3, 31},
    }

    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := maxSum(tt.A, tt.l, tt.m)
            if out != tt.exp {
                t.Fatalf("with input A:%v and l:%d, m:%d wanted %d but got %d",
                    tt.A, tt.l, tt.m, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
