package dup

import "testing"

func TestFindDuplicate(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"1 element", []int{1}, -1},
		{"2 elements", []int{1, 1}, 1},
		{"3 elements", []int{1, 2, 2}, 2},
		{"4 elements", []int{2, 2, 2, 2}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FindDuplicate(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestFindDuplicate2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"1 element", []int{1}, -1},
		{"2 elements", []int{1, 1}, 1},
		{"3 elements", []int{1, 2, 2}, 2},
		{"4 elements", []int{2, 2, 2, 2}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FindDuplicate2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestFindDuplicate3(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"1 element", []int{1}, -1},
		{"2 elements", []int{1, 1}, 1},
		{"3 elements", []int{1, 2, 2}, 2},
		{"4 elements", []int{2, 2, 2, 2}, 2},
		{"8 elements", []int{5, 2, 1, 3, 5, 7, 6, 4}, 5},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FindDuplicate3(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestFindDuplicate4(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"1 element", []int{1}, -1},
		{"2 elements", []int{1, 1}, 1},
		{"3 elements", []int{1, 2, 2}, 2},
		{"4 elements", []int{2, 2, 2, 2}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FindDuplicate4(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
