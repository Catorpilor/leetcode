package up

import "testing"

func TestUniquePath(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name string
		arr  [][]int
		exp  int
	}{
		{"nil slice", nil, 0},
		{"m,n = 1", [][]int{[]int{0, 0, 0}}, 1},
		{"n,m=1 with obstacal", [][]int{[]int{0, 1, 0}}, 0},
		{"faield with m=3,n=2", [][]int{[]int{0, 0}, []int{1, 1}, []int{0, 0}}, 0},
		{"m=3,n=3", [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := UniquePath(c.arr)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.arr)
			}
		})
	}
}

func TestUniquePath2(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name string
		arr  [][]int
		exp  int
	}{
		{"nil slice", nil, 0},
		{"m,n = 1", [][]int{[]int{0, 0, 0}}, 1},
		{"n,m=1 with obstacal", [][]int{[]int{0, 1, 0}}, 0},
		{"faield with m=3,n=2", [][]int{[]int{0, 0}, []int{1, 1}, []int{0, 0}}, 0},
		{"m=3,n=3", [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := UniquePath2(c.arr)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.arr)
			}
		})
	}
}
