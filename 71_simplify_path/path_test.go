package path

import "testing"

func TestSimplifyPath(t *testing.T) {
	st := []struct {
		name string
		path string
		exp  string
	}{
		{"empty", "", ""},
		{"testcase1", "/home/", "/home"},
		{"testcase2", "/../", "/"},
		{"testcase3", "/home//foo/", "/home/foo"},
		{"testcase4", "/a/./b/../../c/", "/c"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := simplifyPath(tt.path)
			if out != tt.exp {
				t.Fatalf("wanted %s but got %s", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
