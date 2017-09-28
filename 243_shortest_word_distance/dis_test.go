package distance

import "testing"

func TestShortestDistance(t *testing.T) {
	st := []struct {
		name  string
		words []string
		w, s  string
		exp   int
	}{
		{"with dups", []string{"practice", "makes", "prefect", "coding", "makes"}, "makes", "coding", 1},
		{"normal", []string{"practice", "makes", "prefect", "coding", "makes"}, "practice", "coding", 3},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := ShortestDistance(c.words, c.w, c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %s and %s",
					c.exp, ret, c.w, c.s)
			}
		})
	}
}
