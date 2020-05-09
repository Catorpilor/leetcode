package ks

import "testing"

func TestKthSmallest(t *testing.T) {
    st := []struct {
        name   string
        matrix [][]int
        k      int
        exp    int
    }{
        {"single row", [][]int{[]int{1, 2, 3}}, 2, 2},
        {"single column", [][]int{[]int{1}, []int{1}, []int{1}}, 1, 3},
        {"testcase1", [][]int{[]int{1, 3, 11}, []int{2, 4, 6}}, 5, 7},
        {"testcase2", [][]int{[]int{1, 3, 11}, []int{2, 4, 6}}, 9, 17},
        {"testcase3", [][]int{[]int{1, 10, 10}, []int{1, 4, 5}, []int{2, 3, 6}}, 7, 9},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := kthSmallest(tt.matrix, tt.k)
            if out != tt.exp {
                t.Fatalf("with input matrix:%v and k:%d wanted %d but got %d", tt.matrix, tt.k, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
