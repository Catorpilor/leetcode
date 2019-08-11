package inter

import (
	"reflect"
	"testing"
)

func TestIntervalIntersection(t *testing.T) {
	st := []struct {
		name string
		a    [][]int
		b    [][]int
		exp  [][]int
	}{
		{"empty slice", [][]int{}, [][]int{[]int{1, 2}}, [][]int{}},
		{"edge 1", [][]int{[]int{1, 3}}, [][]int{[]int{2, 4}, []int{6, 9}}, [][]int{[]int{2, 3}}},
		{"edge 2", [][]int{[]int{1, 100000}}, [][]int{[]int{2, 4}, []int{6, 9}, []int{11, 100001}},
			[][]int{[]int{2, 4}, []int{6, 9}, []int{11, 100000}}},
		{"test 1", [][]int{[]int{0, 2}, []int{5, 10}, []int{13, 23}, []int{24, 25}}, [][]int{[]int{1, 5}, []int{8, 12}, []int{15, 24}, []int{25, 26}},
			[][]int{[]int{1, 2}, []int{5, 5}, []int{8, 10}, []int{15, 23}, []int{24, 24}, []int{25, 25}}},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := IntervalIntersection(tt.a, tt.b)
			if !reflect.DeepEqual(out, tt.exp) {
				t.Fatalf("with input %v and %v, wanted %v but got %v", tt.a, tt.b, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
