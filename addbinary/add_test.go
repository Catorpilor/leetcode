package addbinary

import "testing"

func TestAddBinary(t *testing.T) {
	cases := []struct {
		name   string
		a      string
		b      string
		expect string
	}{
		{"two 1s", "1", "1", "10"},
		{"two 0s", "0", "0", "0"},
		{"1 and 111111", "1", "111111", "1000000"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := AddBinary(c.a, c.b)
			if ret != c.expect {
				t.Fatalf("expected %s, but got %s, with inputs %s and %s",
					c.expect, ret, c.a, c.b)
			}
		})
	}

	t.Log("pass")
}
