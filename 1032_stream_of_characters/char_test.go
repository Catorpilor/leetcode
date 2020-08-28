package character

import "testing"

func initWithTesting(t *testing.T, words []string) *StreamChecker {
	sc := Constructor(words)
	t.Cleanup(func() {
		sc = nil
	})
	return sc
}

func TestWithWords(t *testing.T) {
	sc := initWithTesting(t, []string{"cd", "f", "kl"})
	st := []struct {
		name   string
		letter byte
		exp    bool
	}{
		{"query a", 'a', false},
		{"query b", 'b', false},
		{"query c", 'c', false},
		{"query d", 'd', true},
		{"query e", 'e', false},
		{"query f", 'f', true},
		{"query g", 'g', false},
		{"query h", 'h', false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := sc.Query(tt.letter)
			if out != tt.exp {
				t.Fatalf("when query %c wanted %t but got %t", tt.letter, tt.exp, out)
			}
			t.Log("pass")
		})
	}

}
