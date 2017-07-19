package excel

import "testing"

func TestConvertToTitle(t *testing.T) {
	cases := []struct {
		name   string
		n      int
		expect string
	}{
		{"1", 1, "A"},
		{"27", 27, "AA"},
		{"28", 28, "AB"},
		{"112", 112, "AF"},
		{"1122", 1112, "AF"},
		{"462", 462, "AF"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := ConvertToTitle(c.n)
			if ret != c.expect {
				t.Fatalf("expected %s, but got %s, with input %d",
					c.expect, ret, c.n)
			}
		})
	}
	t.Log("pass")
}
