package peak

import "testing"

func TestPeakIdx(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"single element", []int{1}, 0},
		{"two elements", []int{1, 2}, 1},
		{"test 1", []int{1, 2, 3, 1}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := PeakIdx(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestPeakIdx2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"single element", []int{1}, 0},
		{"two elements", []int{1, 2}, 1},
		{"test 1", []int{1, 2, 3, 1}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := PeakIdx2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestPeakIdx3(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"single element", []int{1}, 0},
		{"two elements", []int{1, 2}, 1},
		{"test 1", []int{1, 2, 3, 1}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := PeakIdx3(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
