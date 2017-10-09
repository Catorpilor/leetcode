package reverse

import "testing"

func TestReverseWords(t *testing.T) {
	st := []struct {
		name, s, exp string
	}{
		{"empty string", "", ""},
		{"single word", "hello", "olleh"},
		{"example", "Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := ReverseWords(c.s)
			if ret != c.exp {
				t.Fatalf("expected %s but got %s, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
