package sum

import (
	"reflect"
	"testing"
)

func TestTwoSums(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		target int
		exp    []int
	}{
		{"empty slice", nil, 0, nil},
		{"not found", []int{1, 2, 3, 4}, 9, nil},
		{"testcase1", []int{2, 11, 15, 7}, 9, []int{0, 3}},
		{"testcase2", []int{20, 3, -11, 5}, 9, []int{0, 2}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := twoSums(tt.nums, tt.target)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input nums: %v and target:%d wanted %v but got %v", tt.nums, tt.target, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
