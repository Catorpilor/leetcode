package sum

import "testing"

func TestNumSubSeqs(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		target int
		exp    int
	}{
		{"testcase1", []int{1, 1, 1, 1}, 2, 15},
		{"testcase2", []int{3, 5, 6, 7}, 9, 4},
		{"testcase3", []int{3, 3, 6, 8}, 10, 6},
		{"testcase4", []int{2, 3, 3, 4, 6, 7}, 12, 61},
		{"testcase5", []int{5, 2, 4, 1, 7, 6, 8}, 16, 127},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := numSubSeq(tt.nums, tt.target)
			if out != tt.exp {
				t.Fatalf("with input nums:%v and target:%d wanted %d but got %d",
					tt.nums, tt.target, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
