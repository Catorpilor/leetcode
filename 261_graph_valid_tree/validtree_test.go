package validtree

import "testing"

func TestValidTree(t *testing.T) {
	st := []struct {
		name  string
		n     int
		edges [][]int
		exp   bool
	}{
		{"0", 0, nil, false},
		{"1", 1, nil, true},
		{"no edges", 4, nil, false},
		{"test1", 5, [][]int{[]int{0, 1}, []int{0, 2}, []int{0, 3}, []int{1, 4}}, true},
		{"test2", 5, [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 3}, []int{1, 3}, []int{1, 4}}, false},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := ValidTree(c.n, c.edges)
			if ret != c.exp {
				t.Fatalf("expected %t but  got %t with input %d and %v",
					c.exp, ret, c.n, c.edges)
			}
		})
	}
}

func TestValidTree2(t *testing.T) {
	st := []struct {
		name  string
		n     int
		edges [][]int
		exp   bool
	}{
		{"0", 0, nil, false},
		{"1", 1, nil, true},
		{"no edges", 4, nil, false},
		{"test1", 5, [][]int{[]int{0, 1}, []int{0, 2}, []int{0, 3}, []int{1, 4}}, true},
		{"test2", 5, [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 3}, []int{1, 3}, []int{1, 4}}, false},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := ValidTree2(c.n, c.edges)
			if ret != c.exp {
				t.Fatalf("expected %t but  got %t with input %d and %v",
					c.exp, ret, c.n, c.edges)
			}
		})
	}
}
