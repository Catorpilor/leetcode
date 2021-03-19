package btf

import (
	"testing"
)


func TestNumberOfFactorBinaryTrees(t *testing.T){
	st := []struct{
		name string
		arr []int
		exp int
	}{
		{"testcase1", []int{2,4},3},
		{"testcase2", []int{2,4,5,10},7},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T){
			out := numberOfFactorBinaryTree(tt.arr)
			if out != tt.exp {
				t.Fatalf("wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
