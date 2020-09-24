package pooling

import "testing"

func TestCarPooling(t *testing.T) {
	st := []struct {
		name     string
		trips    [][]int
		capacity int
		exp      bool
	}{
		{"testcase1", [][]int{[]int{2, 1, 5}, []int{3, 3, 7}}, 4, false},
		{"testcase2", [][]int{[]int{2, 1, 5}, []int{3, 3, 7}}, 5, true},
		{"testcase3", [][]int{[]int{2, 1, 5}, []int{3, 5, 7}}, 3, true},
		{"testcase4", [][]int{[]int{2, 2, 6}, []int{2, 4, 7}, []int{8, 6, 7}}, 11, true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := carPooling(tt.trips, tt.capacity)
			if out != tt.exp {
				t.Fatalf("with trips:%v and capacity:%d wanted %t but got %t", tt.trips, tt.capacity, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
