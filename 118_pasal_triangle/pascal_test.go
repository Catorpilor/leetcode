package pascal

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	cases := []struct {
		name    string
		numRows int
		expect  [][]int
	}{
		{"0 numRows", 0, nil},
		{"1 row", 1, [][]int{[]int{1}}},
		{"4 row", 4, [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}, []int{1, 3, 3, 1}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Generate(c.numRows)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expectd %v, but got %v, with input %d",
					c.expect, ret, c.numRows)
			}
		})
	}
}
