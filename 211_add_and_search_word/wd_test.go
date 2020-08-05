package wd

import "testing"

func initWithTesting(t *testing.T) *WordDict {
	w := Constructor()
	t.Cleanup(func() {
		w = nil
	})
	return w
}

func TestAddWord(t *testing.T) {
	w := initWithTesting(t)
	st := []struct {
		name string
		word string
		exp  bool
	}{
		{"after add hello, it should exsits", "hello", true},
		{"after add hi, it should exists", "hi", true},
		{"after add acb, it should exsits", "acb", true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			w.AddWord(tt.word)
			if w.Search(tt.word) != tt.exp {
				t.Fatalf("with word:%s wanted %t but got %t", tt.word, tt.exp, !tt.exp)
			}
			t.Log("pass")
		})
	}
}

func TestSearch(t *testing.T) {
	w := initWithTesting(t)
	words := []string{"abc", "hello", "hi", "test"}
	for _, word := range words {
		w.AddWord(word)
	}
	st := []struct {
		name string
		word string
		exp  bool
	}{
		{"testcase1", "hello", true},
		{"testcase2", "hi", true},
		{"testcase3", "h....", true},
		{"testcaes4", "h..", false},
		{"testcase5", "ab.", true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := w.Search(tt.word)
			if out != tt.exp {
				t.Fatalf("with input word:%s wanted %t but got %t", tt.word, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
