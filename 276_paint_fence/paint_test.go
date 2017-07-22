package paint

import "testing"

func TestNumWays(t *testing.T) {
	cases := []struct {
		name   string
		n      int
		k      int
		expect int
	}{
		{"n=0", 0, 1, 0},
		{"k=1, n<=2", 2, 1, 1},
		{"k=1, n > 2", 3, 1, 0},
		{"n=k>1", 2, 2, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := NumWays(c.n, c.k)
			if ret != c.expect {
				t.Fatalf("expect %d, but got %d, with input %d and %d",
					c.expect, ret, c.n, c.k)
			}
		})
	}
	t.Log("pass")
}
