package seq

import "testing"

func TestLongestConsecutive(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single", []int{11}, 1},
		{"two non-consecutive", []int{1, 3}, 1},
		{"test 1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"test 2", []int{100, 4, 200, 4, 1, 3, 1, 3, 2}, 4},
		{"identical", []int{1, 1, 1, 1, 1}, 1},
		{"failed 1", []int{0, -1}, 2},
		{"failde 2", []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := LongestConsecutive(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but get %d with input %v", c.exp, ret, c.nums)
			}
		})
	}
}

func TestLongestConsecutive2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single", []int{11}, 1},
		{"two non-consecutive", []int{1, 3}, 1},
		{"test 1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"test 2", []int{100, 4, 200, 4, 1, 3, 1, 3, 2}, 4},
		{"identical", []int{1, 1, 1, 1, 1}, 1},
		{"failed 1", []int{0, -1}, 2},
		{"failde 2", []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := LongestConsecutive2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but get %d with input %v", c.exp, ret, c.nums)
			}
		})
	}
}

func TestLongestConsecutive3(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single", []int{11}, 1},
		{"two non-consecutive", []int{1, 3}, 1},
		{"test 1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"test 2", []int{100, 4, 200, 4, 1, 3, 1, 3, 2}, 4},
		{"identical", []int{1, 1, 1, 1, 1}, 1},
		{"failed 1", []int{0, -1}, 2},
		{"failde 2", []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := LongestConsecutive3(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but get %d with input %v", c.exp, ret, c.nums)
			}
		})
	}
}

func TestLongestConsecutive4(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single", []int{11}, 1},
		{"two non-consecutive", []int{1, 3}, 1},
		{"test 1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"test 2", []int{100, 4, 200, 4, 1, 3, 1, 3, 2}, 4},
		{"identical", []int{1, 1, 1, 1, 1}, 1},
		{"failed 1", []int{0, -1}, 2},
		{"failde 2", []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := LongestConsecutive4(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but get %d with input %v", c.exp, ret, c.nums)
			}
		})
	}
}
