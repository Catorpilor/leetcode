package rand10

import "testing"

func TestRand10(t *testing.T) {
	op := func(a int) bool {
		return a >= 1 && a <= 10
	}
	st := []struct {
		name string
		exp  func(int) bool
	}{
		{"testcase1", op},
		{"testcase2", op},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := rand10()
			if !op(out) {
				t.Fatalf("wanted in the range [1,10], but got %d", out)
			}
			t.Log("pass")
		})
	}
}
