package mad

import "testing"

func TestMinDeleteSum(t *testing.T) {
	st := []struct {
		name   string
		s1, s2 string
		exp    int
	}{
		{"both empty", "", "", 0},
		{"s1 is empty", "", "s", 115},
		{"s1 eq s2", "a", "a", 0},
		{"testcase1", "sea", "eat", 231},
		{"testcase2", "delete", "leet", 403},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minDeleteSum(tt.s1, tt.s2)
			if out != tt.exp {
				t.Fatalf("with s1: %s and s2: %s wanted %d but got %d", tt.s1, tt.s2, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
