package ifc

import "testing"

func setup(chars string, cl int) *CombinationIterator {
	var ci *CombinationIterator
	if ci == nil {
		ci = NewCombineIterator(chars, cl)
	}
	return ci
}

func TestNext(t *testing.T) {
	st := []struct {
		name          string
		characters    string
		combineLength int
		exp           string
	}{
		{"next should be ab", "abc", 2, "ab"},
		{"next should be ac", "abc", 2, "ac"},
		{"next should be bc", "abc", 2, "bc"},
	}
	ci := setup(st[0].characters, st[0].combineLength)
	for _, tt := range st {
		out := ci.Next()
		if out != tt.exp {
			t.Fatalf("name: %s, wanted %s but got %s", tt.name, tt.exp, out)
		}
	}
}

func TestHasNext(t *testing.T) {
	st := []struct {
		name          string
		characters    string
		combineLength int
		exp           bool
	}{
		{"has next should be true", "abc", 2, true},
		{"has next should be true", "abc", 2, true},
		{"has next should be false", "abc", 2, true},
		{"has next should be false 2", "abc", 2, false},
	}
	ci := setup(st[0].characters, st[0].combineLength)
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := ci.HasNext()
			if out != tt.exp {
				t.Fatalf("%s, hasNext should be %t but got %t", tt.name, tt.exp, out)
			}
			ci.Next()
			t.Log("pass")
		})
	}
}
