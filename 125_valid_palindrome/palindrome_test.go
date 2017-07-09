package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect bool
	}{
		{"empty string", "", true},
		{"1 character", "a", true},
		{"2 character", "0P", false},
		{"long", "A man, a plan, a canal: Panama", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := IsPalindrome(c.input)
			if ret != c.expect {
				t.Fatalf("expect %t but got %t, with input %s",
					c.expect, ret, c.input)
			}
		})
	}

	t.Log("pass")
}
