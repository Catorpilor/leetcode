package maze

import "testing"

func TestHasPath(t *testing.T) {
	st := []struct {
		name        string
		maze        [][]int
		start, dest []int
		exp         int
	}{
		{"test1", [][]int{[]int{0, 0, 1, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 1, 0},
			[]int{1, 1, 0, 1, 1}, []int{0, 0, 0, 0, 0}}, []int{0, 4}, []int{4, 4}, 12},
		{"test2", [][]int{[]int{0, 0, 1, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 1, 0},
			[]int{1, 1, 0, 1, 1}, []int{0, 0, 0, 0, 0}}, []int{0, 4}, []int{3, 2}, -1},
		{"test3", [][]int{[]int{0, 0}}, []int{0, 0}, []int{0, 1}, 1},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := HasPath(c.maze, c.start, c.dest)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v and %v and %v",
					c.exp, ret, c.maze, c.start, c.dest)
			}
		})
	}
}

func TestHasPath2(t *testing.T) {
	st := []struct {
		name        string
		maze        [][]int
		start, dest []int
		exp         int
	}{
		{"test1", [][]int{[]int{0, 0, 1, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 1, 0},
			[]int{1, 1, 0, 1, 1}, []int{0, 0, 0, 0, 0}}, []int{0, 4}, []int{4, 4}, 12},
		{"test2", [][]int{[]int{0, 0, 1, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 1, 0},
			[]int{1, 1, 0, 1, 1}, []int{0, 0, 0, 0, 0}}, []int{0, 4}, []int{3, 2}, -1},
		{"test3", [][]int{[]int{0, 0}}, []int{0, 0}, []int{0, 1}, 1},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := HasPath2(c.maze, c.start, c.dest)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v and %v and %v",
					c.exp, ret, c.maze, c.start, c.dest)
			}
		})
	}
}
