package pf

import "testing"

func TestIsPowerOfFour(t *testing.T) {
	st := []struct {
		name string
		num  int
		exp  bool
	}{
		{"num is 1", 1, true},
		{"num is 0", 0, false},
		{"num < 0", -1213, false},
		{"num is 5", 5, false},
		{"num is 16", 16, true},
		{"num is 8", 8, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isPowerOfFour(tt.num)
			if out != tt.exp {
				t.Fatalf("with input num: %d wanted %t but got %t", tt.num, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
