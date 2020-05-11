package flood

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	st := []struct {
		name           string
		image          [][]int
		x, y, newColor int
		exp            [][]int
	}{
		{"empty image", nil, 0, 0, 2, nil},
		{"singe point", [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, 1, 1, 2, [][]int{[]int{0, 0, 0}, []int{0, 2, 0}, []int{0, 0, 0}}},
		{"testcase1", [][]int{[]int{1, 1, 0}, []int{1, 1, 1}, []int{1, 0, 1}}, 1, 1, 2, [][]int{[]int{2, 2, 0}, []int{2, 2, 2}, []int{2, 0, 2}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := floodFill(tt.image, tt.x, tt.y, tt.newColor)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input image:%v and x:%d, y:%d and color:%d wanted %v but got %v", tt.image, tt.x, tt.y, tt.newColor, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
