package gray

import (
	"reflect"
	"testing"
)

func TestGrayCode(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  []int
	}{
		{"edge n=0", 0, []int{0}},
		{"n=2", 2, []int{0, 1, 3, 2}},
		{"n=1", 1, []int{0, 1}},
		{"n=3", 3, []int{0, 1, 3, 2, 6, 7, 5, 4}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := grayCode(tt.n)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input n: %d wanted %v but got %v", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
