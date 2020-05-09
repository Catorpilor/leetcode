package sqrt

import "testing"

func TestIsPerfectSqure(t *testing.T) {
	st := []struct {
		name string
		num  int
		exp  bool
	}{
		{"1", 1, true},
		{"2", 2, false},
		{"4", 4, true},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useMath(c.num)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %d",
					c.exp, ret, c.num)
			}
		})
	}
}

func TestUseBinarySearch(t *testing.T) {
	st := []struct {
		name string
		num  int
		exp  bool
	}{
		{"1", 1, true},
		{"2", 2, false},
		{"4", 4, true},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useBinarySearch(c.num)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %d",
					c.exp, ret, c.num)
			}
		})
	}
}
