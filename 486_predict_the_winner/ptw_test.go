package ptw

import "testing"

func TestPredictTheWinner(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  bool
	}{
		{"single", []int{1}, true},
		{"testcase1", []int{1, 5, 2}, false},
		{"testcase2", []int{1, 5, 233, 7}, true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := predictTheWinner(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums: %v wanted %t but got %t", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestDynamic(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  bool
	}{
		{"single", []int{1}, true},
		{"testcase1", []int{1, 5, 2}, false},
		{"testcase2", []int{1, 5, 233, 7}, true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := dynamic(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums: %v wanted %t but got %t", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
