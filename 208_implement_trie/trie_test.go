package trie

import "testing"

var gtrie *Trie

func prepare() func() {
	if gtrie == nil {
		gtrie = &Trie{}
	}
	return func() {
		gtrie = nil
	}
}

func TestInsert(t *testing.T) {
	f := prepare()
	defer f()
	gtrie.Insert("apple")
	if gtrie.Search("apple") != true {
		t.Fatal("trie inserted apple should exists")
	}
	gtrie.Insert("banana")
	if gtrie.Search("bana") != false {
		t.Fatal("bana shuold not exists")
	}
}

func TestStartWith(t *testing.T) {
	f := prepare()
	defer f()
	gtrie.Insert("apple")
	gtrie.Insert("appoly")
	if gtrie.StartWith("app") != true {
		t.Fatal("startWith app should return true")
	}
}
