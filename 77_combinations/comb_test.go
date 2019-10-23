package comb

import (
	"reflect"
	"testing"
)

func TestCombinations(t *testing.T) {
	st := []struct {
		name string
		n, k int
		exp  [][]int
	}{
		{"invalid case n < k", 2, 4, [][]int{}},
		{"edgecase n = k", 2, 2, [][]int{[]int{1, 2}}},
		{"testcase1", 4, 2, [][]int{[]int{1, 2}, []int{1, 3}, []int{1, 4}, []int{2, 3}, []int{2, 4}, []int{3, 4}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := combinations(tt.n, tt.k)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input n: %d and k: %d, wanted %v but got %v", tt.n, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
