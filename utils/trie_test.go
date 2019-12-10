package utils

import (
	"reflect"
	"testing"
)

var tst *TernarySearchTrie

func setup() {
	tst = NewTST()
	ss := []string{"hello", "hell", "hi", "news", "newspaper"}
	for _, s := range ss {
		tst.Put(s, s)
	}
}

func TestGet(t *testing.T) {
	st := []struct {
		name string
		key  string
		exp  string
	}{
		{"get hello", "hello", "hello"},
		{"get hell", "hell", "hell"},
		{"get his", "his", ""},
	}
	setup()
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := tst.Get(tt.key)
			var v string
			var ok bool
			if v, ok = out.(string); !ok {
				if out != nil {
					t.Fatalf("wanted int but got %v with type: %T", out, out)
				}
			}
			if v != tt.exp {
				t.Fatalf("with key: %s wanted %s but got %s", tt.key, tt.exp, v)
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

func TestWithPrefix(t *testing.T) {
	st := []struct {
		name   string
		prefix string
		exp    []interface{}
	}{
		{"prefix is empty", "", []interface{}{}},
		{"prefix is h", "h", []interface{}{"hello", "hell", "hi"}},
		{"prefix is he", "he", []interface{}{"hello", "hell"}},
		{"prefix is abc", "abc", []interface{}{}},
	}
	setup()
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := tst.WithPrefix(tt.prefix)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with prefix: %s wanted %v but got %v", tt.prefix, tt.exp, out)
			}
			t.Log("pass")
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
		{"query is his", "his", "hi"},
		{"query is newser", "newser", "news"},
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
