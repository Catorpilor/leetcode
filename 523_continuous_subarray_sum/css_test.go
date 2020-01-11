package css

import "testing"

func TestCheckSubArraySum(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  bool
	}{
		{"empty slice", []int{}, 2, false},
		{"not found", []int{1, 2, 3}, 4, false},
		{"testcase1", []int{1, 1, 1}, 2, true},
		{"testcase2", []int{23, 2, 4, 6, 7}, 6, true},
		{"failed1", []int{0, 0}, 0, true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := checkSubArraySum(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("with nums: %v and k: %d wanted %t but got %t", tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}

}

func TestPrefixSum(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  bool
	}{
		{"empty slice", []int{}, 2, false},
		{"not found", []int{1, 2, 3}, 4, false},
		{"testcase1", []int{1, 1, 1}, 2, true},
		{"testcase2", []int{23, 2, 4, 6, 7}, 6, true},
		{"failed1", []int{0, 0}, 0, true},
		{"failed2", []int{0, 1, 0}, 0, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := preSumHash(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("with nums: %v and k: %d wanted %t but got %t", tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}

}
