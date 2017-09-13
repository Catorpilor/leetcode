package chain

import "testing"

func TestFindLongestChain(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{"nil slice", nil, 0},
		{"one element", [][]int{[]int{1, 2}}, 1},
		{"two element", [][]int{[]int{1, 2}, []int{2, 3}}, 1},
		{"three element", [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}}, 2},
		{"three element", [][]int{[]int{3, 4}, []int{2, 3}, []int{1, 2}}, 2},
		{"randome elelment", [][]int{[]int{-10, -8}, []int{8, 9}, []int{-5, 0}, []int{6, 10}, []int{-6, -4}, []int{1, 7}, []int{9, 10}, []int{-4, 7}}, 4},
		{"randome elelment", [][]int{[]int{-6, 9}, []int{1, 6}, []int{8, 10}, []int{-1, 4}, []int{-6, -2}, []int{-9, 8}, []int{-5, 3}, []int{0, 3}}, 3},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := FindLongestChain(c.input)
			if ret != c.expect {
				t.Fatalf("expect %d but got %d, with input %v",
					c.expect, ret, c.input)
			}
		})
	}
	t.Log("pass")
}
