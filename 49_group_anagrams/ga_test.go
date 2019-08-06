package ga

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	st := []struct {
		name  string
		input []string
		exp   [][]string
	}{
		{"nil input", nil, [][]string{}},
		{"empty input", []string{}, [][]string{}},
		{"input's length eq 1", []string{"test"}, [][]string{[]string{"test"}}},
		{"input's length eq 2 with anagrams", []string{"test", "tste"}, [][]string{[]string{"test", "tste"}}},
		{"leetcode test1", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{
			[]string{"eat", "tea", "ate"},
			[]string{"tan", "nat"},
			[]string{"bat"},
		}},
	}

	for _, s := range st {
		t.Run(s.name, func(t *testing.T) {
			out := GroupAnagrams(s.input)
			if !reflect.DeepEqual(out, s.exp) {
				t.Fatalf("input %v wanted %v but got %v", s.input, s.exp, out)
			}
			t.Log("pass")
		})
	}
}
