package utils

import "testing"

func TestNumOfOnes(t *testing.T) {
	st := []struct {
		name   string
		a, exp int
	}{
		{"a=0", 0, 0},
		{"a=1", 1, 1},
		{"a=8", 8, 1},
		{"a=7", 7, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := NumOfOnes(tt.a)
			if out != tt.exp {
				t.Fatalf("with input a:%d wanted %d but got %d", tt.a, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
