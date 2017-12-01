package task

import "testing"

func TestLeastInterval(t *testing.T) {
	st := []struct {
		name   string
		tasks  []byte
		n, exp int
	}{
		{"empty slice", []byte{}, 2, 0},
		{"single element", []byte{'A'}, 3, 1},
		{"identicals", []byte{'A', 'A', 'A'}, 3, 9},
		{"a3b3", []byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := LeastInterval(c.tasks, c.n)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v and %d",
					c.exp, ret, c.tasks, c.n)
			}
		})
	}
}
