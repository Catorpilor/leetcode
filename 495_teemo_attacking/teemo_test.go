package teemo

import "testing"

func TestTeemoAttack(t *testing.T) {
	st := []struct {
		name       string
		timeSeries []int
		duration   int
		exp        int
	}{
		{"testcase1", []int{1, 4}, 2, 4},
		{"testcase2", []int{1, 2}, 2, 3},
		{"testcase3", []int{1, 2, 3}, 2, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := teemoAttack(tt.timeSeries, tt.duration)
			if out != tt.exp {
				t.Fatalf("with input timeSeries: %v and duration:%d wanted %d but got %d", tt.timeSeries, tt.duration, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
