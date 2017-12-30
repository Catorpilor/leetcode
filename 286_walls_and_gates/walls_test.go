package walls

import (
	"reflect"
	"testing"
)

// const INF = math.MaxInt32

func TestWallsAndGates(t *testing.T) {
	st := []struct {
		name  string
		rooms [][]int
		exp   [][]int
	}{
		{"empty", [][]int{}, [][]int{}},
		{"single inf", [][]int{[]int{INF}}, [][]int{[]int{INF}}},
		{"test1", [][]int{[]int{INF, -1, 0, INF}, []int{INF, INF, INF, -1}, []int{INF, -1, INF, -1},
			[]int{0, -1, INF, INF}}, [][]int{[]int{3, -1, 0, 1}, []int{2, 2, 1, -1}, []int{1, -1, 2, -1}, []int{0, -1, 3, 4}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := WallsAndGates(c.rooms)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.rooms)
			}
		})
	}
}

func TestWallsAndGates2(t *testing.T) {
	st := []struct {
		name  string
		rooms [][]int
		exp   [][]int
	}{
		{"empty", [][]int{}, [][]int{}},
		{"single inf", [][]int{[]int{INF}}, [][]int{[]int{INF}}},
		{"test1", [][]int{[]int{INF, -1, 0, INF}, []int{INF, INF, INF, -1}, []int{INF, -1, INF, -1},
			[]int{0, -1, INF, INF}}, [][]int{[]int{3, -1, 0, 1}, []int{2, 2, 1, -1}, []int{1, -1, 2, -1}, []int{0, -1, 3, 4}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := WallsAndGates2(c.rooms)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.rooms)
			}
		})
	}
}
