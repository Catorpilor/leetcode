package limit

import "testing"

func TestLongestSubarray(t *testing.T) {
    st := []struct {
        name  string
        nums  []int
        limit int
        exp   int
    }{
        {"single num", []int{1}, 5, 1},
        {"all identical", []int{1, 1, 1, 1}, 2, 4},
        {"testcase1", []int{8, 2, 4, 7}, 4, 2},
        {"testcase2", []int{10, 1, 2, 4, 7, 2}, 5, 4},
        {"testcase3", []int{4, 2, 2, 2, 4, 4, 2, 2}, 0, 3},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := longestSubarray(tt.nums, tt.limit)
            if out != tt.exp {
                t.Fatalf("with input nums:%v and limit:%d wanted %d but got %d", tt.nums, tt.limit, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestLongestSubarrayUseDeque(t *testing.T) {
    st := []struct {
        name  string
        nums  []int
        limit int
        exp   int
    }{
        {"single num", []int{1}, 5, 1},
        {"all identical", []int{1, 1, 1, 1}, 2, 4},
        {"testcase1", []int{8, 2, 4, 7}, 4, 2},
        {"testcase2", []int{10, 1, 2, 4, 7, 2}, 5, 4},
        {"testcase3", []int{4, 2, 2, 2, 4, 4, 2, 2}, 0, 3},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := useDeque(tt.nums, tt.limit)
            if out != tt.exp {
                t.Fatalf("with input nums:%v and limit:%d wanted %d but got %d", tt.nums, tt.limit, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
