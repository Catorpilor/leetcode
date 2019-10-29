package ps

import "testing"

func TestGetPermutation(t *testing.T) {
	st := []struct {
		name string
		n, k int
		exp  string
	}{
		{"n=1,k=1", 1, 1, "1"},
		{"n=2,k=2", 2, 2, "21"},
		{"n=3,k=3", 3, 3, "213"},
		{"n=4,k=9", 4, 9, "2314"},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := getPermutation(tt.n, tt.k)
			if out != tt.exp {
				t.Fatalf("with input n: %d and k: %d wanted %s but got %s", tt.n, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
