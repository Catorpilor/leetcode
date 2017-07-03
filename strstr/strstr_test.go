package strstr

import "testing"

func TestStrStr(t *testing.T) {
	cases := []struct {
		name     string
		haystack string
		needle   string
		expect   int
	}{
		{"empty haystack and needle", "", "", 0},
		{"empty haystack", "", "1", -1},
		{"hello_el", "hello", "el", 1},
		{"aaa_aaaa", "aaa", "aaaa", -1},
		{"aaa_bbb", "aaa", "bbb", -1},
		{"mississippi_pi", "mississippi", "pi", 9},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := StrStr(c.haystack, c.needle)
			if ret != c.expect {
				t.Fatalf("expected %d, but got %d, with inputs (%s,%s)",
					c.expect, ret, c.haystack, c.needle)
			}
		})
	}
	t.Log("passed")
}
