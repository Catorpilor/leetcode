package excel

import "testing"

func TestTitleToNumber(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect int
	}{
		{"just A", "A", 1},
		{"just Z", "Z", 26},
		{"AB", "AB", 28},
		{"APT", "APT", 1112},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := TitleToNumber(c.input)
			if ret != c.expect {
				t.Fatalf("expected %d, but got %d, with inputs %s",
					c.expect, ret, c.input)
			}
		})
	}
	t.Log("passed")
}
