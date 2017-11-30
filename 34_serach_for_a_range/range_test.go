package ranger

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		target int
		exp    []int
	}{
		{"empty slice", []int{}, 2, []int{-1, -1}},
		{"identical", []int{1, 1, 1, 1}, 1, []int{0, 3}},
		{"notfound", []int{1, 2, 3}, 4, []int{-1, -1}},
		{"test 1", []int{1}, 1, []int{0, 0}},
		{"faield 1", []int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := SearchRange(c.nums, c.target)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v and %d",
					c.exp, ret, c.nums, c.target)
			}
		})
	}
}

func TestSearchRange2(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		target int
		exp    []int
	}{
		{"empty slice", []int{}, 2, []int{-1, -1}},
		{"identical", []int{1, 1, 1, 1}, 1, []int{0, 3}},
		{"notfound", []int{1, 2, 3}, 4, []int{-1, -1}},
		{"test 1", []int{1}, 1, []int{0, 0}},
		{"faield 1", []int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := SearchRange2(c.nums, c.target)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v and %d",
					c.exp, ret, c.nums, c.target)
			}
		})
	}
}
