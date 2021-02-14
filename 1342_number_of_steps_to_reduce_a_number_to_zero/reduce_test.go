package reduce

import "testing"

func TestNumOfSteps(t *testing.T) {
	st := []struct {
		name string
		num  int
		exp  int
	}{
		{"num=0", 0, 0},
		{"testcase1", 8, 4},
		{"testcase2", 14, 6},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := numOfSteps(tt.nums)
			if out != tt.exp {
				t.Fatalf("wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
