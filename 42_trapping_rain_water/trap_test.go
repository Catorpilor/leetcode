package trap

import "testing"

func TestTrap(t *testing.T) {
	st := []struct {
		name   string
		height []int
		exp    int
	}{
		{"empty", []int{}, 0},
		{"single", []int{1}, 0},
		{"two", []int{1, 3}, 0},
		{"identical", []int{1, 1, 1, 1}, 0},
		{"test 1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Trap(c.height)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.height)
			}
		})
	}
}

func TestTrap2(t *testing.T) {
	st := []struct {
		name   string
		height []int
		exp    int
	}{
		{"empty", []int{}, 0},
		{"single", []int{1}, 0},
		{"two", []int{1, 3}, 0},
		{"identical", []int{1, 1, 1, 1}, 0},
		{"test 1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Trap2(c.height)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.height)
			}
		})
	}
}

func TestBF(t *testing.T) {
	st := []struct {
		name   string
		height []int
		exp    int
	}{
		{"empty", []int{}, 0},
		{"single", []int{1}, 0},
		{"two", []int{1, 3}, 0},
		{"identical", []int{1, 1, 1, 1}, 0},
		{"test 1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := bruteForce(c.height)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.height)
			}
		})
	}
}
