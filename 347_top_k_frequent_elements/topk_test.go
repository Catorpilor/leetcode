package topk

import (
	"reflect"
	"testing"
)

func TestTopKFreq(t *testing.T) {
	st := []struct {
		name  string
		input []int
		k     int
		exp   []int
	}{
		{"empty slice", []int{}, 0, []int{}},
		{"single element", []int{1}, 1, []int{1}},
		{"all distinct", []int{1, 2, 3, 4, 5, 6, 7}, 7, []int{7, 1, 2, 3, 4, 5, 6}},
		{"leetcode test case1", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := TopKFreq(tt.input, tt.k)
			if !reflect.DeepEqual(out, tt.exp) {
				t.Fatalf("with input %v and k: %d, wanted %v but got %v",
					tt.input, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
