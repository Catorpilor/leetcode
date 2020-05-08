package check

import "testing"

func TestIsStraight(t *testing.T) {
    st := []struct {
        name  string
        coors [][]int
        exp   bool
    }{
        {"empty coors", nil, false},
        {"one point", [][]int{[]int{1, 2}}, false},
        {"two points", [][]int{[]int{1, 2}, []int{3, 4}}, true},
        {"horizontal", [][]int{[]int{1, 2}, []int{2, 2}, []int{3, 2}, []int{4, 2}}, true},
        {"vertical", [][]int{[]int{3, 0}, []int{3, 1}, []int{3, 2}, []int{3, 3}}, true},
        {"testcase1", [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{4, 5}, []int{5, 6}}, true},
        {"testcase2", [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 5}, []int{4, 5}, []int{5, 6}}, false},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := isStraight(tt.coors)
            if out != tt.exp {
                t.Fatalf("with input coors: %v wanted %t but got %t", tt.coors, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
