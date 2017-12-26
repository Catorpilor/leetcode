package flatten

import (
	"math"
	"reflect"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestFlatten(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"empty", []int{}, nil},
		{"single", []int{1}, []int{1}},
		{"test 1", []int{1, 2, 3}, []int{1, 2, 3}},
		{"test2", []int{1, 2, 5, 3, 4, math.MinInt32, 6}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := Flatten(root)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestFlatten2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"empty", []int{}, nil},
		{"single", []int{1}, []int{1}},
		{"test 1", []int{1, 2, 3}, []int{1, 2, 3}},
		{"test2", []int{1, 2, 5, 3, 4, math.MinInt32, 6}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := Flatten2(root)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestFlatte3n(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"empty", []int{}, nil},
		{"single", []int{1}, []int{1}},
		{"test 1", []int{1, 2, 3}, []int{1, 2, 3}},
		{"test2", []int{1, 2, 5, 3, 4, math.MinInt32, 6}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := Flatten3(root)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
