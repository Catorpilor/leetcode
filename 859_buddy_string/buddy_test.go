package buddy

import "testing"

func TestCanBeABuddyString(t *testing.T) {
	st := []struct {
		name string
		a, b string
		exp  bool
	}{
		{"testcase1", "ab", "ba", true},
		{"testcase2", "ab", "ab", false},
		{"testcase3", "aa", "aa", true},
		{"testcase4", "aaaaaaabc", "aaaaaaacb", true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := buddyString(tt.a, tt.b)
			if out != tt.exp {
				t.Fatalf("with input a:%s and b:%s wanted %t but got %t", tt.a, tt.b, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
