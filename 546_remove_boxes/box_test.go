package box

import "testing"

func TestRemoveBoxes(t *testing.T) {
	st := []struct {
		name  string
		boxes []int
		exp   int
	}{
		{"empty", nil, 0},
		{"single", []int{1}, 1},
		{"all identical", []int{1, 1, 1}, 9},
		{"test 1", []int{1, 3, 2, 2, 2, 3, 4, 3, 1}, 23},
		{"test2", []int{1, 4, 1, 4, 1}, 11},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := RemoveBoxes(c.boxes)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.boxes)
			}
		})
	}
}
