package meeting

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	st := []struct {
		name      string
		intervals []Interval
		exp       bool
	}{
		{"single interval", []Interval{{1, 2}}, true},
		{"without collapse", []Interval{{25, 30}, {5, 10}, {15, 20}}, true},
		{"with collapse", []Interval{{0, 30}, {5, 10}, {15, 20}}, false},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CanAttendMeetings(c.intervals)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with inputs %v",
					c.exp, ret, c.intervals)
			}
		})
	}
}
