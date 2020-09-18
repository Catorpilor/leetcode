package image

import "testing"

func TestLargestOverlap(t *testing.T) {
	st := []struct {
		name string
		a, b [][]int
		exp  int
	}{
		{"testcase1", [][]int{[]int{1}}, [][]int{[]int{1}}, 1},
		{"testcase2", [][]int{[]int{0}}, [][]int{[]int{0}}, 0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := largestOverlap(tt.a, tt.b)
			if out != tt.exp {
				t.Fatalf("with input a:%v, b:%v wanted %d but got %d", tt.a, tt.b, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
