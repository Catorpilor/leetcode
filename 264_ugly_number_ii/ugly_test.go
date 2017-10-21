package ugly

import "testing"

func TestNthUglyNumber(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name   string
		n, exp int
	}{
		{"n eq 1", 1, 1},
		{"n eq 3", 3, 3},
		{"n eq 10", 10, 12},
		{"n eq 14", 14, 20},
		{"n eq 394", 394, 291600},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := NthUglyNumber(c.n)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}

func TestNthUglyNumber2(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name   string
		n, exp int
	}{
		{"n eq 1", 1, 1},
		{"n eq 3", 3, 3},
		{"n eq 10", 10, 12},
		{"n eq 14", 14, 20},
		{"n eq 394", 394, 291600},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := NthUglyNumber2(c.n)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}
