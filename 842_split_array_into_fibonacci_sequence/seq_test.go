package seq

import (
	"reflect"
	"testing"
)

func TestSplitIntoFib(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  []int
	}{
		{"len(s)=1", "2", []int{}},
		{"len(s)=2", "12", []int{}},
		{"valid len(s)=3", "123", []int{1, 2, 3}},
		{"invalid with leading 0", "011", []int{0, 1, 1}},
		{"valid with 0", "101", []int{1, 0, 1}},
		{"testcase1", "112358130", []int{}},
		{"testcase2", "11235813", []int{1, 1, 2, 3, 5, 8, 13}},
		{"failed1", "539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511", []int{}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := splitIntoFib(tt.s)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input s:%s wanted %v but got %v", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}

}
