package min

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		s, exp int
	}{
		{"empty slice", []int{}, 5, 0},
		{"single element less than s", []int{1}, 2, 0},
		{"test 1", []int{2, 3, 1, 2, 4, 3}, 7, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MinSubArrayLen(c.s, c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %d and %v",
					c.exp, ret, c.s, c.nums)
			}
		})
	}
}

func TestMinSubArrayLen2(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		s, exp int
	}{
		{"empty slice", []int{}, 5, 0},
		{"single element less than s", []int{1}, 2, 0},
		{"test 1", []int{2, 3, 1, 2, 4, 3}, 7, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MinSubArrayLen2(c.s, c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %d and %v",
					c.exp, ret, c.s, c.nums)
			}
		})
	}
}

func TestMinSubArrayLen3(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		s, exp int
	}{
		{"empty slice", []int{}, 5, 0},
		{"single element less than s", []int{1}, 2, 0},
		{"test 1", []int{2, 3, 1, 2, 4, 3}, 7, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MinSubArrayLen3(c.s, c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %d and %v",
					c.exp, ret, c.s, c.nums)
			}
		})
	}
}
