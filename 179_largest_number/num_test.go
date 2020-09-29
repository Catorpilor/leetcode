package num

import "testing"

func TestLargestNum(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  string
	}{
		{"testcase1", []int{0, 0}, "0"},
		{"testcase2", []int{1}, "1"},
		{"testcase3", []int{10, 2}, "210"},
		{"testcase4", []int{3, 30, 34, 5, 9}, "934339"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := largestNum(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums:%v wanted %s but got %s", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
