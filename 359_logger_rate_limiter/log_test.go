package log

import "testing"

var lg Logger = Constructor()

func TestShouldPrintMessage(t *testing.T) {
	st := []struct {
		name    string
		message string
		ts      int
		exp     bool
	}{
		{"foo at 1", "foo", 1, true},
		{"bar at 1", "bar", 1, true},
		{"zoo at 1", "zoo", 1, true},
		{"foo at 5", "foo", 5, false},
		{"foo at 11", "foo", 11, true},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := lg.ShouldPrintMessage(c.message, c.ts)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %s and %d",
					c.exp, ret, c.message, c.ts)
			}
		})
	}
}
