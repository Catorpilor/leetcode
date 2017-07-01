package strstr

import "testing"

type res struct {
	haystack string
	needle   string
	expect   int
}

func TestStrStr(t *testing.T) {
	cases := []*res{
		&res{"", "", 0},
		&res{"", "1", -1},
		&res{"hello", "el", 1},
		&res{"aaa", "aaaa", -1},
		&res{"aaa", "bbb", -1},
		&res{"mississippi", "pi", 9},
	}
	for _, c := range cases {
		ret := StrStr(c.haystack, c.needle)
		if ret != c.expect {
			t.Fatalf("expected %d, but got %d, with inputs (%s,%s)",
				c.expect, ret, c.haystack, c.needle)
		}
	}
	t.Log("passed")
}
