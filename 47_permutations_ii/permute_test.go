package permute

import "testing"

func TestPermuteUnique(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  [][]int
	}{
		{"empty nums", []int{}, [][]int{}},
		{"len eq 1", []int{1}, [][]int{[]int{1}}},
		{"all identical", []int{1, 1}, [][]int{[]int{1, 1}}},
		{"testcase1", []int{1, 2, 3},
			[][]int{[]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3},
				[]int{3, 2, 1}, []int{3, 1, 2}, []int{2, 3, 1}}},
		{"failed1", []int{2, 2, 1, 1}, [][]int{[]int{1, 1, 2, 2}, []int{1, 2, 1, 2},
			[]int{1, 2, 2, 1}, []int{2, 1, 1, 2}, []int{2, 1, 2, 1}, []int{2, 2, 1, 1}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := permuteUnique(tt.nums)
			if !deepEqual(tt.exp, out) {
				t.Fatalf("with input nums: %v wanted %v but got %v", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestPermuteUnique2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  [][]int
	}{
		{"empty nums", []int{}, [][]int{}},
		{"len eq 1", []int{1}, [][]int{[]int{1}}},
		{"all identical", []int{1, 1}, [][]int{[]int{1, 1}}},
		{"testcase1", []int{1, 2, 3},
			[][]int{[]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3},
				[]int{3, 2, 1}, []int{3, 1, 2}, []int{2, 3, 1}}},
		{"failed1", []int{2, 2, 1, 1}, [][]int{[]int{1, 1, 2, 2}, []int{1, 2, 1, 2},
			[]int{1, 2, 2, 1}, []int{2, 1, 1, 2}, []int{2, 1, 2, 1}, []int{2, 2, 1, 1}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := permuteUnique2(tt.nums)
			if !deepEqual(tt.exp, out) {
				t.Fatalf("with input nums: %v wanted %v but got %v", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func deepEqual(arr1, arr2 [][]int) bool {
	n1, n2 := len(arr1), len(arr2)
	if n1 != n2 {
		return false
	}
	hset := make(map[int]int)
	for i := range arr1 {
		for j := range arr1[i] {
			hset[arr1[i][j]]++
		}
	}
	for i := range arr2 {
		for j := range arr2[i] {
			hset[arr2[i][j]]--
		}
	}

	for _, v := range hset {
		if v != 0 {
			return false
		}
	}
	return true
}
