package rotate

import "testing"

func TestSearch(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		target int
		exp    bool
	}{
		{"empty slice", []int{}, 1, false},
		{"single elment", []int{1}, 2, false},
		{"test1", []int{4, 5, 6, 7, 7, 7, 8, 0, 1, 2, 3}, 7, true},
		{"failed 1", []int{1, 3, 1, 1, 1}, 3, true},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Search(c.nums, c.target)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t with input %v and %d",
					c.exp, ret, c.nums, c.target)
			}
		})
	}
}

func TestSearch2(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		target int
		exp    bool
	}{
		{"empty slice", []int{}, 1, false},
		{"single elment", []int{1}, 2, false},
		{"test1", []int{4, 5, 6, 7, 7, 7, 8, 0, 1, 2, 3}, 7, true},
		{"failed 1", []int{1, 3, 1, 1, 1}, 3, true},
		{"failde 2", []int{3, 1}, 1, true},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Search2(c.nums, c.target)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t with input %v and %d",
					c.exp, ret, c.nums, c.target)
			}
		})
	}
}

func TestSearch3(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		target int
		exp    bool
	}{
		{"empty slice", []int{}, 1, false},
		{"single elment", []int{1}, 2, false},
		{"test1", []int{4, 5, 6, 7, 7, 7, 8, 0, 1, 2, 3}, 7, true},
		{"failed 1", []int{1, 3, 1, 1, 1}, 3, true},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Search3(c.nums, c.target)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t with input %v and %d",
					c.exp, ret, c.nums, c.target)
			}
		})
	}
}
