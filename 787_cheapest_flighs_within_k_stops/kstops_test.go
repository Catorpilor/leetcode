package kstops

import "testing"

func TestCheapestFlight(t *testing.T) {
	st := []struct {
		name              string
		n                 int
		flights           [][]int
		src, dest, k, exp int
	}{
		{"testcase1", 3, [][]int{[]int{0, 1, 100}, []int{1, 2, 100}, []int{0, 2, 500}}, 0, 2, 1, 200},
		{"testcase2", 3, [][]int{[]int{0, 1, 100}, []int{1, 2, 100}, []int{0, 2, 500}}, 0, 2, 0, 500},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := findCheapest(tt.n, tt.flights, tt.src, tt.dest, tt.k)
			if out != tt.exp {
				t.Fatalf("with input n:%d, flights:%v, src:%d, dest:%d, and k:%d wanted %d but got %d", tt.n, tt.flights,
					tt.src, tt.dest, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
