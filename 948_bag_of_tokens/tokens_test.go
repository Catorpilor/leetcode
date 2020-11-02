package tokens

import "testing"

func TestMaxScore(t *testing.T) {
	st := []struct {
		name   string
		tokens []int
		p      int
		exp    int
	}{
		{"testcase1", []int{100}, 50, 0},
		{"testcase2", []int{25}, 50, 1},
		{"testcase3", []int{100, 200}, 150, 1},
		{"testcase4", []int{400, 300, 200, 100}, 200, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxScore(tt.tokens, tt.p)
			if out != tt.exp {
				t.Fatalf("with input tokens:%v and p:%d wanted %d but got %d", tt.tokens, tt.p, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
