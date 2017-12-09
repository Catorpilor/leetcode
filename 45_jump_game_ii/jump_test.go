package jump

import "testing"

func TestJump(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single", []int{1}, 0},
		{"two", []int{1, 2}, 1},
		{"all identical", []int{1, 1, 1, 1}, 3},
		{"test 1", []int{2, 3, 1, 1, 4}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Jump(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v", c.exp, ret, c.nums)
			}
		})
	}
}

func TestJump2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single", []int{1}, 0},
		{"two", []int{1, 2}, 1},
		{"all identical", []int{1, 1, 1, 1}, 3},
		{"test 1", []int{2, 3, 1, 1, 4}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Jump2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v", c.exp, ret, c.nums)
			}
		})
	}
}

func TestJump3(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, -1},
		{"single", []int{1}, 0},
		{"two", []int{1, 2}, 1},
		{"all identical", []int{1, 1, 1, 1}, 3},
		{"test 1", []int{2, 3, 1, 1, 4}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Jump3(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v", c.exp, ret, c.nums)
			}
		})
	}
}
