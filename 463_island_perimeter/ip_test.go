package ip

import "testing"

func TestIslandPerimeter(t *testing.T) {
    st := []struct {
        name string
        grid [][]int
        exp  int
    }{
        {"empty grid", [][]int{}, 0},
        {"testcase1", [][]int{[]int{1, 0, 0, 0}}, 4},
        {"testcase2", [][]int{[]int{1, 1, 0, 0}}, 6},
        {"testcase3", [][]int{[]int{1, 1, 1, 1}}, 10},
        {"testcase3", [][]int{[]int{0, 1, 0, 0}, []int{1, 1, 1, 0}, []int{0, 1, 0, 0}, []int{1, 1, 0, 0}}, 16},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := islandPerimeter(tt.grid)
            if out != tt.exp {
                t.Fatalf("with grid: %v wanted %d but got %d", tt.grid, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestIslandPerimeterWithIter(t *testing.T) {
    st := []struct {
        name string
        grid [][]int
        exp  int
    }{
        {"empty grid", [][]int{}, 0},
        {"testcase1", [][]int{[]int{1, 0, 0, 0}}, 4},
        {"testcase2", [][]int{[]int{1, 1, 0, 0}}, 6},
        {"testcase3", [][]int{[]int{1, 1, 1, 1}}, 10},
        {"testcase3", [][]int{[]int{0, 1, 0, 0}, []int{1, 1, 1, 0}, []int{0, 1, 0, 0}, []int{1, 1, 0, 0}}, 16},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := iterator(tt.grid)
            if out != tt.exp {
                t.Fatalf("with grid: %v wanted %d but got %d", tt.grid, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
