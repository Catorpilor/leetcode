package coins

import "testing"

func TestMaxCoins(t *testing.T) {
	st := []struct {
		name  string
		piles []int
		exp   int
	}{
		{"testcaes1", []int{2, 4, 1, 2, 7, 8}, 9},
		{"testcase2", []int{2, 4, 5}, 4},
		{"testcase3", []int{9, 8, 7, 6, 5, 1, 2, 3, 4}, 18},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxCoins(tt.piles)
			if out != tt.exp {
				t.Fatalf("with input piles:%v wanted %d but got %d", tt.piles, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
