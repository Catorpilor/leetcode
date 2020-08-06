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
	words := []string{"at", "and", "an", "add", "a", "bat"}
	for _, word := range words {
		w.AddWord(word)
	}
	st := []struct {
		name string
		word string
		exp  bool
	}{
		{"testcase1", "hello", false},
		{"testcase2", "hi", false},
		{"testcase3", "h....", false},
		{"testcase4", "h..", false},
		{"testcase5", "ab.", false},
		{"ts6", ".at", true},
		{"ts7", "an.", true},
		{"ts8", "a.d.", false},
		{"ts9", ".", true},
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
