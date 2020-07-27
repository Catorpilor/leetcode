package wb

import "testing"

func TestWatterBottles(t *testing.T) {
	st := []struct {
		name                    string
		numBottles, numExchange int
		exp                     int
	}{
		{"{1,2}", 1, 2, 1},
		{"testcase1", 15, 4, 19},
		{"testcase2", 5, 5, 6},
		{"testcase3", 2, 3, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := watterBottles(tt.numBottles, tt.numExchange)
			if out != tt.exp {
				t.Fatalf("with input numBottles:%d and numExchange:%d wanted %d but got %d", tt.numBottles, tt.numExchange, tt.exp, out)
			}
			t.Log("passef")
		})
	}
}
