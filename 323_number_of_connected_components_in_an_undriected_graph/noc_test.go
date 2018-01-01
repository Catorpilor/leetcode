package noc

import "testing"

func TestCountComponents(t *testing.T) {
	st := []struct {
		name  string
		n     int
		edges [][]int
		exp   int
	}{
		{"n=0", 0, nil, 0},
		{"n=1", 1, nil, 1},
		{"n=2", 2, [][]int{[]int{0, 1}}, 1},
		{"n=5", 5, [][]int{[]int{0, 1}, []int{1, 2}, []int{3, 4}}, 2},
		{"failed1", 5, [][]int{[]int{0, 1}, []int{1, 2}, []int{0, 2}, []int{3, 4}}, 2},
		{"failed2", 5, [][]int{[]int{0, 1}, []int{2, 1}, []int{2, 0}, []int{2, 4}}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CountComponents(c.n, c.edges)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %d and %v",
					c.exp, ret, c.n, c.edges)
			}
		})
	}
}

func TestCountComponents2(t *testing.T) {
	st := []struct {
		name  string
		n     int
		edges [][]int
		exp   int
	}{
		{"n=0", 0, nil, 0},
		{"n=1", 1, nil, 1},
		{"n=2", 2, [][]int{[]int{0, 1}}, 1},
		{"n=5", 5, [][]int{[]int{0, 1}, []int{1, 2}, []int{3, 4}}, 2},
		{"failed1", 5, [][]int{[]int{0, 1}, []int{1, 2}, []int{0, 2}, []int{3, 4}}, 2},
		{"failed2", 5, [][]int{[]int{0, 1}, []int{2, 1}, []int{2, 0}, []int{2, 4}}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CountComponents2(c.n, c.edges)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %d and %v",
					c.exp, ret, c.n, c.edges)
			}
		})
	}
}

func TestCountComponents3(t *testing.T) {
	st := []struct {
		name  string
		n     int
		edges [][]int
		exp   int
	}{
		{"n=0", 0, nil, 0},
		{"n=1", 1, nil, 1},
		{"n=2", 2, [][]int{[]int{0, 1}}, 1},
		{"n=5", 5, [][]int{[]int{0, 1}, []int{1, 2}, []int{3, 4}}, 2},
		{"failed1", 5, [][]int{[]int{0, 1}, []int{1, 2}, []int{0, 2}, []int{3, 4}}, 2},
		{"failed2", 5, [][]int{[]int{0, 1}, []int{2, 1}, []int{2, 0}, []int{2, 4}}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CountComponents3(c.n, c.edges)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %d and %v",
					c.exp, ret, c.n, c.edges)
			}
		})
	}
}
