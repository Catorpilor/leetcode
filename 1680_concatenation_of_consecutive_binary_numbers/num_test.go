package num

import "testing"

func TestConcatenateBinary(t *testing.T) {
	st := []struct{
		name string
		n int
		exp int
	}{
		{"n=1", 1, 1},
		{"n=3", 3, 27},
		{"n=12", 12, 505379714},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := cancatenateBinary(tt.n)
			if out != tt.exp {
				t.Fatalf("wanted %d but got %d", tt.exp, out)
			}
			t.Log(:"pass")
		})
	}
}
