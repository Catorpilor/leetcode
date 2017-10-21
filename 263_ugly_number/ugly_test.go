package ugly

import "testing"

func TestUglyNumber(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name string
		n    int
		exp  bool
	}{
		{"n eq 1", 1, true},
		{"n eq 2", 2, true},
		{"n eq 3", 3, true},
		{"n eq 4", 4, true},
		{"n eq 14", 14, false},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := UglyNumber(c.n)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}
