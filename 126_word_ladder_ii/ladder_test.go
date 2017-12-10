package ladder

import "testing"

func TestLadders(t *testing.T) {
	st := []struct {
		name, start, end string
		wordlists        []string
		exp              [][]string
	}{
		{"test1", "hit", "cog", []string{"dot", "hot", "lot", "dog", "log", "cog"}, [][]string{
			[]string{"hit", "hot", "dot", "dog", "cog"},
			[]string{"hit", "hot", "lot", "log", "cog"},
		}},
		{"tle", "qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb",
			"sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to",
			"ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn",
			"au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge",
			"th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa",
			"pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn",
			"mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr",
			"sq", "ye"}, [][]string{}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Ladders(c.start, c.end, c.wordlists)
			if len(ret) != len(c.exp) {
				t.Fatalf("expected %v but got %v with input %s, %s and %v",
					c.exp, ret, c.start, c.end, c.wordlists)
			}
		})
	}
}
