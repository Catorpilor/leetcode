package str

import "testing"

func TestGetSmallestString(t *testing.T) {
	st := []struct {
		name string
		n    int
		k    int
		exp  string
	}{
		{"n=3,k=27", 3, 27, "aay"},
		{"n=5,k=73", 5, 73, "aaszz"},
		{"n=3,k=8", 3, 8, "aaf"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := getSmallestString(tt.n, tt.k)
			if out != tt.exp {
				t.Fatalf("wanted %s but got %s", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
