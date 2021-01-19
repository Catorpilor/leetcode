package vowel

import "testing"

func TestCountVowelString(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"n=1", 1, 5},
		{"n=2", 2, 15},
		{"n=33", 33, 66045},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := countVowelString(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestCountVowelStringuseDP(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"n=1", 1, 5},
		{"n=2", 2, 15},
		{"n=33", 33, 66045},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useDP(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
