package as

import "testing"

func TestNumberOfArithmeticSlices(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 0},
		{"two elements", []int{1, 2}, 0},
		{"3 elements", []int{1, 2, 3}, 1},
		{"test 13579", []int{1, 3, 5, 7, 9}, 6},
		{"test 11257", []int{1, 1, 2, 5, 7}, 0},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := NumberOfArithmeticSlices(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestNumberOfArithmeticSlices2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 0},
		{"two elements", []int{1, 2}, 0},
		{"3 elements", []int{1, 2, 3}, 1},
		{"test 13579", []int{1, 3, 5, 7, 9}, 6},
		{"test 11257", []int{1, 1, 2, 5, 7}, 0},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := NumberOfArithmeticSlices2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
