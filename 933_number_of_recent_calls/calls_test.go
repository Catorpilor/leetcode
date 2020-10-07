package calls

import "testing"

func constructWithTesting(t *testing.T) *RecentCalls {
	rc := Constructor()
	t.Cleanup(func() {
		rc = nil
	})
	return rc
}

func TestRecentCalls(t *testing.T) {
	st := []struct {
		name string
		t    int
		exp  int
	}{
		{"testcase1", 1, 1},
		{"testcase2", 100, 2},
		{"testcase3", 3001, 3},
		{"testcase4", 3002, 3},
	}
	rc := constructWithTesting(t)
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := rc.Ping(tt.t)
			if out != tt.exp {
				t.Fatalf("with input t:%d wanted %d but got %d", tt.t, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
