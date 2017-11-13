package query

import "testing"

var matrix = [][]int{
	[]int{3, 0, 1, 4, 2},
	[]int{5, 6, 3, 2, 1},
	[]int{1, 2, 0, 1, 5},
	[]int{4, 1, 0, 1, 7},
	[]int{1, 0, 3, 0, 5},
}

var nm = Constructor(matrix)

func TestSumRegion(t *testing.T) {
	st := []struct {
		name                   string
		row1, col1, row2, col2 int
		exp                    int
	}{
		{"2143", 2, 1, 4, 3, 8},
		{"1122", 1, 1, 2, 2, 11},
		{"1224", 1, 2, 2, 4, 12},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := nm.SumRegion(c.row1, c.col1, c.row2, c.col2)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d ,with input %d,%d,%d,%d",
					c.exp, ret, c.row1, c.col1, c.row2, c.col2)
			}
		})
	}
}
