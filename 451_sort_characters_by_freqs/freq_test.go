package freq

import "testing"

func TestFreqSort(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"all identical", "aaaa", "aaaa"},
		{"testcase1", "tree", "eetr"},
		{"testcase2", "cccaaaa", "aaaaccc"},
		{"testcase3", "Aabb", "bbAa"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := freqSort(tt.s)
			if out != tt.exp {
				t.Fatalf("with input s:%s wanted %s but got %s", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
