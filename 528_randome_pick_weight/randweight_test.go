package randweight

import (
	"reflect"
	"testing"
)

func TestRandomPick(t *testing.T){
	st := []struct{
		name string
		arr []int
		numOfPicks int
		exp []int
	}{
		{"single element", []int{1}, 2,  []int{0,0}},
		{"equal weight", []int{2,2}, 2, []int{0,1}},
		{"testcase1", []int{1,3},4, []int{1,0,1,1}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T){
			rw := Constructor(tt.arr)
			res := make([]int, 0, tt.numOfPicks)
			for i := 0; i < tt.numOfPicks; i++ {
				res = append(res, rw.randomPick())
			}
			if !reflect.DeepEqual(tt.exp, res) {
				t.Fatalf("with input arr: %v wanted %v but got %v", tt.arr, tt.exp, res)
			}
		})
	}
}