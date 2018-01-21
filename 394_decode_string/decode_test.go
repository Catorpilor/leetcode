package decode

import "testing"

func TestDecodeString(t *testing.T) {
	st := []struct {
		name, s, exp string
	}{
		{"empty", "", ""},
		{"justk", "2", ""},
		{"invalidform", "2[", ""},
		{"test1", "2[a]", "aa"},
		{"test2", "3[a2[c]]", "accaccacc"},
		{"test3", "2[abc]3[cd]ef", "abcabccdcdcdef"},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := DecodeString(c.s)
			if ret != c.exp {
				t.Fatalf("expected %s but got %s with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}

func TestDecodeString2(t *testing.T) {
	st := []struct {
		name, s, exp string
	}{
		{"empty", "", ""},
		{"justk", "2", ""},
		{"invalidform", "2[", ""},
		{"test1", "2[a]", "aa"},
		{"test2", "3[a2[c]]", "accaccacc"},
		{"test3", "2[abc]3[cd]ef", "abcabccdcdcdef"},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := DecodeString2(c.s)
			if ret != c.exp {
				t.Fatalf("expected %s but got %s with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
