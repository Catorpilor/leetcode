package ascii

import "testing"

func TestMinDelAscii(t *testing.T) {
	st := []struct {
		name, s1, s2 string
		exp          int
	}{
		{"both is empty", "", "", 0},
		{"one is empty", "", "abc", 294},
		{"no common sequence", "a", "b", 195},
		{"ab ba", "ab", "ba", 194},
		{"test1", "delete", "leet", 403},
		{"test2", "sea", "eat", 231},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MinDelAscii(c.s1, c.s2)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %s and %s",
					c.exp, ret, c.s1, c.s2)
			}
		})
	}
}

func TestMinDelAscii2(t *testing.T) {
	st := []struct {
		name, s1, s2 string
		exp          int
	}{
		{"both is empty", "", "", 0},
		{"one is empty", "", "abc", 294},
		{"no common sequence", "a", "b", 195},
		{"ab ba", "ab", "ba", 194},
		{"test1", "delete", "leet", 403},
		{"test2", "sea", "eat", 231},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MinDelAscii2(c.s1, c.s2)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %s and %s",
					c.exp, ret, c.s1, c.s2)
			}
		})
	}
}
