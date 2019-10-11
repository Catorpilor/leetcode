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

func TestInorderTraversal(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"empty", nil, nil},
		{"single", []int{1}, []int{1}},
		{"two nodes", []int{1, 2}, []int{2, 1}},
		{"test 1", []int{1, 2, 3}, []int{2, 1, 3}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := ConstructTree(c.nums)
			ret := InorderTraversal(root)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestPreorderTraversal(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"empty", nil, nil},
		{"single", []int{1}, []int{1}},
		{"two nodes", []int{1, 2}, []int{1, 2}},
		{"test 1", []int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := ConstructTree(c.nums)
			ret := PreorderTraversal(root)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestPostorderTraversal(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"empty", nil, nil},
		{"single", []int{1}, []int{1}},
		{"two nodes", []int{1, 2}, []int{2, 1}},
		{"test 1", []int{1, 2, 3}, []int{2, 3, 1}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := ConstructTree(c.nums)
			ret := PostorderTraversal(root)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestIsEqual(t *testing.T) {
	st := []struct {
		name string
		l, r *TreeNode
		exp  bool
	}{
		{"both empty", nil, nil, true},
		{"one is empty", ConstructTree([]int{1}), nil, false},
		{"equal", ConstructTree([]int{4, 1, 7, 3, 2}), ConstructTree([]int{4, 1, 7, 3, 2}), true},
		{"single node", ConstructTree([]int{4}), ConstructTree([]int{4}), true},
		{"testcase1", ConstructTree([]int{4, 1, 7, 3, 2}), ConstructTree([]int{4, 1, 7, 3, 2, 5}), false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := IsEqual(tt.l, tt.r)
			if out != tt.exp {
				t.Fatalf("with input l:%v and r:%v wanted %t but got %t", InorderTraversal(tt.l), InorderTraversal(tt.r), tt.exp, out)
			}
		})
		t.Log("pass")
	}
}
