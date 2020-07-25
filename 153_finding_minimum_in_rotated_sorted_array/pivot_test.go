package pivot

import "testing"

func TestFindMinUseStraight(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty slice", []int{}, 0},
		{"single", []int{1}, 1},
		{"test 1", []int{4, 5, 6, 7, 0, 1, 2}, 0},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useStraight(c.nums)
			if ret != c.exp {
				t.Fatalf("expectde %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestFindMinUseBS(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty slice", []int{}, 0},
		{"single", []int{1}, 1},
		{"test 1", []int{4, 5, 6, 7, 0, 1, 2}, 0},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useBinarySearch(c.nums)
			if ret != c.exp {
				t.Fatalf("expectde %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestFindMinUseBS2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty slice", []int{}, 0},
		{"single", []int{1}, 1},
		{"test 1", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"testcase2", []int{4, 4, 4, 4, 5, 6, 7, 7, 7, 1, 1, 1, 2, 3}, 1},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useBS(c.nums)
			if ret != c.exp {
				t.Fatalf("expectde %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}

}
