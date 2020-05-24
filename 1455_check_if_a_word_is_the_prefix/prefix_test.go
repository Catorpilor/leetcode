package prefix

import "testing"

func TestIsPrefix(t *testing.T) {
	st := []struct {
		name     string
		sentence string
		word     string
		exp      int
	}{
		{"single word", "hello", "world", -1},
		{"testcase1", "i am tired", "you", -1},
		{"testcase2", "i use tripple pillow", "pill", 4},
		{"testcase3", "hello from the other side", "ot", 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isPrefix(tt.sentence, tt.word)
			if out != tt.exp {
				t.Fatalf("with input sentence:%s and word:%s wanted %d but got %d", tt.sentence, tt.word, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
