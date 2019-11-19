package utils

import "testing"

var rtrie *RTrie

func rsetup() {
	if rtrie == nil {
		rtrie = NewRTrie(26)
	}
	words := []string{"hello", "hell", "hi", "journal", "job", "news", "newspaper"}
	for _, word := range words {
		rtrie.Put(word)
	}
}

func TestRtrieContains(t *testing.T) {
	st := []struct {
		name  string
		query string
		exp   bool
	}{
		{"contains hello?", "hello", true},
		{"contains hell?", "hell", true},
		{"contains hi?", "hi", true},
		{"contains helloer?", "helloer", false},
	}
	rsetup()
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := rtrie.Contains(tt.query)
			if out != tt.exp {
				t.Fatalf("with query:%s wanted %t but got %t", tt.query, tt.exp, out)
			}
		})
	}
}

func TestRtrieGet(t *testing.T) {
	st := []struct {
		name  string
		query string
		exp   interface{}
	}{
		{"get hello?", "hello", "hello"},
		{"get hell?", "hell", "hell"},
		{"get hi?", "hi", "hi"},
		{"get helloer?", "helloer", nil},
	}
	rsetup()
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := rtrie.Get(tt.query)
			if out != tt.exp {
				t.Fatalf("with query:%s wanted %v but got %v", tt.query, tt.exp, out)
			}
		})
	}
}
