package angle

import "testing"

func TestAngleClock(t *testing.T) {
	st := []struct {
		name          string
		hour, minutes int
		exp           float64
	}{
		{"00:00", 0, 0, 0.0},
		{"12:30", 12, 30, 165},
		{"3:30", 3, 30, 75},
		{"3:15", 3, 15, 7.5},
		{"12:00", 12, 0, 0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := angleClock(tt.hour, tt.minutes)
			if out != tt.exp {
				t.Fatalf("with input hour:%d and minutes:%d wanted %f but got %f", tt.hour, tt.minutes, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
