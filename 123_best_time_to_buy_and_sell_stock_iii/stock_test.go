package stock

import "testing"

func TestStock(t *testing.T) {
	st := []struct {
		name   string
		prices []int
		exp    int
	}{
		{"empty", []int{}, 0},
		{"single", []int{}, 0},
		{"two decreasing", []int{2, 1}, 0},
		{"test 1", []int{7, 1, 5, 3, 6, 4}, 7},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Stock(c.prices)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.prices)
			}
		})
	}
}

func TestStock2(t *testing.T) {
	st := []struct {
		name   string
		prices []int
		exp    int
	}{
		{"empty", []int{}, 0},
		{"single", []int{}, 0},
		{"two decreasing", []int{2, 1}, 0},
		{"test 1", []int{7, 1, 5, 3, 6, 4}, 7},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Stock2(c.prices)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.prices)
			}
		})
	}
}
