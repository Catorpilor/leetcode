package judge

import "testing"

func TestFindJudge(t *testing.T) {
    st := []struct {
        name  string
        n     int
        trust [][]int
        exp   int
    }{
        {"single", 1, nil, 1},
        {"two person", 2, [][]int{[]int{1, 2}}, 2},
        {"testcase1", 3, [][]int{[]int{1, 3}, []int{2, 3}}, 3},
        {"testcase2", 3, [][]int{[]int{1, 3}, []int{2, 3}, []int{3, 1}}, -1},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := findJudge(tt.n, tt.trust)
            if out != tt.exp {
                t.Fatalf("with input n:%d and trust:%v wanted %d but got %d", tt.n, tt.trust, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
