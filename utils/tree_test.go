package utils

import (
	"math"
	"reflect"
	"testing"
)

func TestContstructTree(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"empty slice", []int{}, nil},
		{"single", []int{1}, []int{1}},
		{"test1", []int{1, 2, 3, math.MinInt32, 4}, []int{1, 2, 3, 4}},
		{"test2", []int{1, math.MinInt32, 2, 3}, []int{1, 2, 3}},
		{"test4", []int{1, 2, 3, 4, math.MinInt32, 5,
			math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32, 6},
			[]int{1, 2, 3, 4, 5, 6}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := ConstructTree(c.nums)
			ret := LevelOrderTravesal(root)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
