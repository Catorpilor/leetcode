package pos

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	st := []struct {
		name   string
		arr    []int
		target int
		exp    []int
	}{
		{"all identical, target is 2", []int{2, 2, 2, 2, 2, 2}, 2, []int{0, 5}},
		{"all identical, target not found", []int{1, 1, 1, 1, 1}, 2, []int{-1, -1}},
		{"testcase1", []int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{"testcase2", []int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := searchRange(tt.arr, tt.target)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input arr: %v and target:%d wanted %v but got %v", tt.arr, tt.target, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
