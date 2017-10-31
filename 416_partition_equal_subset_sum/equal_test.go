package equal

import "testing"

func TestCanPartition(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  bool
	}{
		{"single element", []int{1}, false},
		{"two identical", []int{2, 2}, true},
		{"two different", []int{1, 2}, false},
		{"test 15115", []int{1, 5, 11, 5}, true},
		{"test 1235", []int{1, 2, 3, 5}, false},
		{"failed 1111", []int{1, 1, 1, 1}, true},
		{"failed 22", []int{66, 90, 7, 6, 32, 16, 2, 78, 69, 88, 85, 26, 3, 9, 58, 65, 30, 96, 11, 31, 99, 49, 63, 83, 79, 97, 20, 64, 81, 80, 25, 69, 9, 75, 23, 70, 26, 71, 25, 54, 1, 40, 41, 82, 32, 10, 26, 33, 50, 71, 5, 91, 59, 96, 9, 15, 46, 70, 26, 32, 49, 35, 80, 21, 34, 95, 51, 66, 17, 71, 28, 88, 46, 21, 31, 71, 42, 2, 98, 96, 40, 65, 92, 43, 68, 14, 98, 38, 13, 77, 14, 13, 60, 79, 52, 46, 9, 13, 25, 8}, true},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CanPartition(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestCanPartition2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  bool
	}{
		{"single element", []int{1}, false},
		{"two identical", []int{2, 2}, true},
		{"two different", []int{1, 2}, false},
		{"test 15115", []int{1, 5, 11, 5}, true},
		{"test 1235", []int{1, 2, 3, 5}, false},
		{"failed 1111", []int{1, 1, 1, 1}, true},
		{"failed 22", []int{66, 90, 7, 6, 32, 16, 2, 78, 69, 88, 85, 26, 3, 9, 58, 65, 30, 96, 11, 31, 99, 49, 63, 83, 79, 97, 20, 64, 81, 80, 25, 69, 9, 75, 23, 70, 26, 71, 25, 54, 1, 40, 41, 82, 32, 10, 26, 33, 50, 71, 5, 91, 59, 96, 9, 15, 46, 70, 26, 32, 49, 35, 80, 21, 34, 95, 51, 66, 17, 71, 28, 88, 46, 21, 31, 71, 42, 2, 98, 96, 40, 65, 92, 43, 68, 14, 98, 38, 13, 77, 14, 13, 60, 79, 52, 46, 9, 13, 25, 8}, true},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CanPartition2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
