package cb

import (
	"reflect"
	"testing"
)

func TestAssignBikes(t *testing.T) {
	st := []struct {
		name           string
		workers, bikes [][]int
		exp            []int
	}{
		{"empty workers", [][]int{}, [][]int{[]int{1, 2}}, []int{}},
		{"testcase1", [][]int{[]int{0, 0}, []int{2, 1}}, [][]int{[]int{1, 2}, []int{3, 3}}, []int{1, 0}},
		{"testcase2", [][]int{[]int{0, 0}, []int{1, 1}, []int{2, 0}}, [][]int{[]int{1, 0}, []int{2, 2}, []int{2, 1}}, []int{0, 2, 1}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := assignBikes(tt.workers, tt.bikes)
			if !reflect.DeepEqual(out, tt.exp) {
				t.Fatalf("with workers %v and bikes %v wanted %v but got %v", tt.workers, tt.bikes, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
