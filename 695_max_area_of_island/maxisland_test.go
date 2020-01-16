package maxisland

import "testing"

func TestMaxArea(t *testing.T) {
    st := []struct {
        name string
        grid [][]int
        exp  int
    }{
        {"empty grid", [][]int{}, 0},
        {"all zeros", [][]int{[]int{0, 0, 0, 0}}, 0},
        {"all ones", [][]int{[]int{1, 1, 1, 1}}, 4},
        {"testcases1", [][]int{[]int{1, 0}, []int{0, 1}}, 1},
        {"testcase2", [][]int{[]int{0, 1, 0}, []int{1, 1, 1}, []int{0, 1, 1}}, 6},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := maxArea(tt.grid)
            if out != tt.exp {
                t.Fatalf("with grid: %v wanted %d but got %d", tt.grid, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
