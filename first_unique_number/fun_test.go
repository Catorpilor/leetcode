package fun

import "testing"

func TestShowFirstUnique(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			fu := FU{}
			out := fu.showFirstUnique()
			if out != tt.exp {
				t.Fatalf("with input nums: %v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
