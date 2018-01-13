package balloon

import "testing"

func TestMaxCoins(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", nil, 0},
		{"single", []int{1}, 1},
		{"two", []int{2, 3}, 9},
		{"identical", []int{2, 2, 2}, 14},
		{"test1", []int{3, 1, 5, 8}, 167},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxCoins(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestMaxCoins2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", nil, 0},
		{"single", []int{1}, 1},
		{"two", []int{2, 3}, 9},
		{"identical", []int{2, 2, 2}, 14},
		{"test1", []int{3, 1, 5, 8}, 167},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxCoins2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
