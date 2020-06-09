package coin

import "testing"

func TestChange(t *testing.T) {
	st := []struct {
		name   string
		coins  []int
		amount int
		exp    int
	}{
		{"only one coin", []int{1}, 5, 1},
		{"testcase1", []int{1, 2, 3}, 5, 5},
		{"testcase2", []int{2}, 3, 0},
		{"testcase3", []int{10}, 10, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := change(tt.amount, tt.coins)
			if out != tt.exp {
				t.Fatalf("with input amount:%d and coins:%v wanted %d but got %d", tt.amount, tt.coins, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
func TestChangeUseBf(t *testing.T) {
	st := []struct {
		name   string
		coins  []int
		amount int
		exp    int
	}{
		{"only one coin", []int{1}, 5, 1},
		{"testcase1", []int{1, 2, 3}, 5, 5},
		{"testcase2", []int{2}, 3, 0},
		{"testcase3", []int{10}, 10, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useBF(tt.amount, tt.coins)
			if out != tt.exp {
				t.Fatalf("with input amount:%d and coins:%v wanted %d but got %d", tt.amount, tt.coins, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
