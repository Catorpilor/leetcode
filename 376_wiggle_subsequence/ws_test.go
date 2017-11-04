package ws

import "testing"

func TestWiggleMaxLength(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single slice", []int{1}, 1},
		{"two elements", []int{1, 2}, 2},
		{"test 1", []int{1, 7, 4, 9, 2, 5}, 6},
		{"test 2", []int{1, 17, 10, 13, 15, 10, 5, 16, 8}, 7},
		{"incresing", []int{1, 3, 5, 7, 9}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := WiggleMaxLength(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d ,with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestWiggleMaxLengthDP1(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single slice", []int{1}, 1},
		{"two elements", []int{1, 2}, 2},
		{"test 1", []int{1, 7, 4, 9, 2, 5}, 6},
		{"test 2", []int{1, 17, 10, 13, 15, 10, 5, 16, 8}, 7},
		{"incresing", []int{1, 3, 5, 7, 9}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := WiggleMaxLengthDP1(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d ,with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestWiggleMaxLengthDP2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single slice", []int{1}, 1},
		{"two elements", []int{1, 2}, 2},
		{"test 1", []int{1, 7, 4, 9, 2, 5}, 6},
		{"test 2", []int{1, 17, 10, 13, 15, 10, 5, 16, 8}, 7},
		{"incresing", []int{1, 3, 5, 7, 9}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := WiggleMaxLengthDP2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d ,with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestWiggleMaxLengthDP3(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single slice", []int{1}, 1},
		{"two elements", []int{1, 2}, 2},
		{"test 1", []int{1, 7, 4, 9, 2, 5}, 6},
		{"test 2", []int{1, 17, 10, 13, 15, 10, 5, 16, 8}, 7},
		{"incresing", []int{1, 3, 5, 7, 9}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := WiggleMaxLengthDP3(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d ,with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestWiggleMaxLengthGreedy(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single slice", []int{1}, 1},
		{"two elements", []int{1, 2}, 2},
		{"test 1", []int{1, 7, 4, 9, 2, 5}, 6},
		{"test 2", []int{1, 17, 10, 13, 15, 10, 5, 16, 8}, 7},
		{"incresing", []int{1, 3, 5, 7, 9}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := WiggleMaxLengthGreedy(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d ,with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
