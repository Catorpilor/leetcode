package ms

import "testing"

func TestMaxSquare(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name    string
		maxtrix [][]int
		exp     int
	}{
		{"nil matrix", nil, 0},
		{"1 row or one column with 1", [][]int{[]int{0, 1, 0}}, 1},
		{"1 row or one column without 1", [][]int{[]int{0}, []int{0}, []int{0}}, 0},
		{"test cases", [][]int{[]int{1, 0, 1, 0, 0}, []int{1, 0, 1, 1, 1}, []int{1, 1, 1, 1, 1}, []int{1, 0, 0, 1, 0}}, 4},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := maxSquare(c.maxtrix)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.maxtrix)
			}
		})
	}
}

func TestMaxSquareUseDP(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name    string
		maxtrix [][]int
		exp     int
	}{
		{"nil matrix", nil, 0},
		{"non nil matrix", [][]int{[]int{}}, 0},
		{"1 row or one column with 1", [][]int{[]int{0, 1, 0}}, 1},
		{"1 row or one column without 1", [][]int{[]int{0}, []int{0}, []int{0}}, 0},
		{"test cases", [][]int{[]int{1, 0, 1, 0, 0}, []int{1, 0, 1, 1, 1}, []int{1, 1, 1, 1, 1}, []int{1, 0, 0, 1, 0}}, 4},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useDP(c.maxtrix)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.maxtrix)
			}
		})
	}
}

func TestMaxSquareUseDPWithLessSpace(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name    string
		maxtrix [][]int
		exp     int
	}{
		{"nil matrix", nil, 0},
		{"non nil matrix", [][]int{[]int{}}, 0},
		{"1 row or one column with 1", [][]int{[]int{0, 1, 0}}, 1},
		{"1 row or one column without 1", [][]int{[]int{0}, []int{0}, []int{0}}, 0},
		{"test cases", [][]int{[]int{1, 0, 1, 0, 0}, []int{1, 0, 1, 1, 1}, []int{1, 1, 1, 1, 1}, []int{1, 0, 0, 1, 0}}, 4},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useDPWithLessSpace(c.maxtrix)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.maxtrix)
			}
		})
	}
}
