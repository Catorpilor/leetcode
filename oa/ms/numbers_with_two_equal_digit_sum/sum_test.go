package sum

import "testing"

func TestDigitSum(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty nums", nil, -1},
		{"no such pari", []int{51, 32, 43}, -1},
		{"all equal", []int{111, 21, 12, 30}, 141},
		{"testcase1", []int{51, 71, 17, 42}, 93},
		{"testcase2", []int{42, 33, 60}, 102},
		{"testcase3", []int{-111, 21, 12, 30, 60}, 51},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := digitSum(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums: %v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
