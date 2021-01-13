package boat

import "testing"

func TestMinRes(t *testing.T) {
	st := []struct {
		name  string
		ppl   []int
		limit int
		exp   int
	}{
		{"testcase1", []int{1, 2}, 3, 1},
		{"testcase2", []int{3, 3, 4, 5}, 5, 4},
		{"testcase3", []int{3, 2, 2, 1}, 3, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minRes(tt.ppl, tt.limit)
			if out != tt.exp {
				t.Fatalf("with input ppl:%v and limit:%d wanted %d but got %d", tt.ppl, tt.limit, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
