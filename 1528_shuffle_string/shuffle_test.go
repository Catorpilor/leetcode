package shuffle

import "testing"

func TestShuffleString(t *testing.T) {
	st := []struct {
		name    string
		s       string
		indices []int
		exp     string
	}{
		{"empty string", "", nil, ""},
		{"testcase1", "codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}, "leetcode"},
		{"testcase2", "abc", []int{0, 1, 2}, "abc"},
		{"testcase3", "aiohn", []int{3, 1, 4, 2, 0}, "nihao"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := shuffleString(tt.s, tt.indices)
			if out != tt.exp {
				t.Fatalf("with input s:%s and indices:%v wanted %s but got %s", tt.s, tt.indices, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
