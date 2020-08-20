package goat

import "testing"

func TestToGoatLatins(t *testing.T) {
	st := []struct {
		name     string
		str, exp string
	}{
		{"testcase1", "I speak Goat Latin", "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"},
		{"testcase2", "The quick brown fox jumped over the lazy dog", "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := toGoatLatin(tt.str)
			if out != tt.exp {
				t.Fatalf("with input str:%s wanted %s but got %s", tt.str, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
