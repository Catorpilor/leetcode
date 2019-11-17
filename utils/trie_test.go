package utils

import (
	"testing"
)

var tst *TernarySearchTrie

func setup() {
	tst = NewTST()
	ss := []string{"hello", "hell", "hi", "news", "newspaper"}
	for _, s := range ss {
		tst.Put(s, 1)
	}
}

func TestGet(t *testing.T) {
	st := []struct {
		name string
		key  string
		exp  int
	}{
		{"get hello", "hello", 1},
		{"get hell", "hell", 1},
		{"get his", "his", 0},
	}
	setup()
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := tst.Get(tt.key)
			var v int
			var ok bool
			if v, ok = out.(int); !ok {
				if out != nil {
					t.Fatalf("wanted int but got %v with type: %T", out, out)
				}
			}
			if v != tt.exp {
				t.Fatalf("with key: %s wanted %d but got %d", tt.key, tt.exp, v)
			}
		})
	}
}

func TestContains(t *testing.T) {
	st := []struct {
		name string
		key  string
		exp  bool
	}{
		{"contains hello", "hello", true},
		{"contains hell", "hell", true},
		{"contains his", "his", false},
	}
	setup()
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := tst.Contains(tt.key)
			if out != tt.exp {
				t.Fatalf("with key: %s wanted %t but got %t", tt.key, tt.exp, out)
			}
		})
	}
}

func TestLongestPrefixOf(t *testing.T) {
	st := []struct {
		name  string
		query string
		exp   string
	}{
		{"query is helloer", "helloer", "hello"},
		{"query is hello", "hello", "hello"},
		{"not exists", "jimmy", ""},
	}
	setup()
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := tst.LongestPrefixOf(tt.query)
			if out != tt.exp {
				t.Fatalf("with query: %s waned %s but got %s", tt.query, tt.exp, out)
			}
		})
	}
}
