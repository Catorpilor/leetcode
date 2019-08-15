package game

import "testing"

func TestJudgePoint(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  bool
	}{
		{"test1", []int{1, 2, 3, 4}, true},
		{"test2", []int{1, 1, 1, 1}, false},
		{"test3", []int{1, 2, 1, 2}, false},
		{"test4", []int{9, 9, 9, 9}, false},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := JudgePoint(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
