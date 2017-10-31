package ts

import "testing"

func TestTargetSum(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		s, exp int
	}{
		{"nil slice", nil, 1, 0},
		{"empty slice", []int{}, 1, 0},
		{"test 12345", []int{1, 2, 3, 4, 5}, 5, 3},
		{"test 11111 and s = 3", []int{1, 1, 1, 1, 1}, 3, 5},
		{"test 2112 and s = 0", []int{2, 1, 1, 2}, 0, 4},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := TargetSum(c.nums, c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v and %d",
					c.exp, ret, c.nums, c.s)
			}
		})
	}
}

func TestTargetSum2(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		s, exp int
	}{
		{"nil slice", nil, 1, 0},
		{"empty slice", []int{}, 1, 0},
		{"test 12345", []int{1, 2, 3, 4, 5}, 5, 3},
		{"test 11111 and s = 3", []int{1, 1, 1, 1, 1}, 3, 5},
		{"test 2112 and s = 0", []int{2, 1, 1, 2}, 0, 4},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := TargetSum2(c.nums, c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v and %d",
					c.exp, ret, c.nums, c.s)
			}
		})
	}
}

func TestTargetSum3(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		s, exp int
	}{
		{"nil slice", nil, 1, 0},
		{"empty slice", []int{}, 1, 0},
		{"test 12345", []int{1, 2, 3, 4, 5}, 5, 3},
		{"test 11111 and s = 3", []int{1, 1, 1, 1, 1}, 3, 5},
		{"test 2112 and s = 0", []int{2, 1, 1, 2}, 0, 4},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := TargetSum3(c.nums, c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v and %d",
					c.exp, ret, c.nums, c.s)
			}
		})
	}
}

func TestTargetSumWithMemo(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		s, exp int
	}{
		{"nil slice", nil, 1, 0},
		{"empty slice", []int{}, 1, 0},
		{"test 12345", []int{1, 2, 3, 4, 5}, 5, 3},
		{"test 11111 and s = 3", []int{1, 1, 1, 1, 1}, 3, 5},
		{"test 2112 and s = 0", []int{2, 1, 1, 2}, 0, 4},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := TargetSumWithMemo(c.nums, c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v and %d",
					c.exp, ret, c.nums, c.s)
			}
		})
	}
}
