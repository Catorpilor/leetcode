package closest

import (
	"reflect"
	"testing"
)

func TestKClosest(t *testing.T) {
	st := []struct {
		name   string
		points [][]int
		k      int
		exp    [][]int
	}{
		{"just one point", [][]int{[]int{3, 4}}, 1, [][]int{[]int{3, 4}}},
		{"testcase1", [][]int{[]int{1, 3}, []int{-2, 2}}, 1, [][]int{[]int{-2, 2}}},
		{"testcase2", [][]int{[]int{3, 3}, []int{5, -1}, []int{-2, 4}}, 2, [][]int{[]int{-2, 4}, []int{3, 3}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := kClosest(tt.points, tt.k)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input points: %v and k:%d wanted %v but got %v", tt.points, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestKClosestQuickSelect(t *testing.T) {
	st := []struct {
		name   string
		points [][]int
		k      int
		exp    [][]int
	}{
		{"just one point", [][]int{[]int{3, 4}}, 1, [][]int{[]int{3, 4}}},
		{"testcase1", [][]int{[]int{1, 3}, []int{-2, 2}}, 1, [][]int{[]int{-2, 2}}},
		{"testcase2", [][]int{[]int{3, 3}, []int{5, -1}, []int{-2, 4}}, 2, [][]int{[]int{-2, 4}, []int{3, 3}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useQuickSelect(tt.points, tt.k)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input points: %v and k:%d wanted %v but got %v", tt.points, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
