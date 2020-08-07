package dup

import "testing"

func TestFindDuplicates(t *testing.T) {
	st := []struct {
		name      string
		nums, exp []int
	}{
		{"empty", []int{}, []int{}},
		{"single element", []int{1}, []int{}},
		{"identical", []int{1, 1}, []int{1}},
		{"test 1", []int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := findDuplicates(c.nums)
			if len(ret) != len(c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestUseConstantSpace(t *testing.T) {
	st := []struct {
		name      string
		nums, exp []int
	}{
		{"empty", []int{}, []int{}},
		{"single element", []int{1}, []int{}},
		{"identical", []int{1, 1}, []int{1}},
		{"test 1", []int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useConstantSpace(c.nums)
			if len(ret) != len(c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
