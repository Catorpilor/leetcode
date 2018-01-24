package course

import "testing"

func TestCanFinish(t *testing.T) {
	st := []struct {
		name       string
		numCourses int
		pres       [][]int
		exp        bool
	}{
		{"zero courses", 0, nil, true},
		{"signle course", 1, nil, true},
		{"no pres", 4, nil, true},
		{"test1", 2, [][]int{[]int{1, 0}}, true},
		{"test4", 2, [][]int{[]int{0, 1}}, true},
		{"test2", 2, [][]int{[]int{1, 0}, []int{0, 1}}, false},
		{"test3", 3, [][]int{[]int{1, 2}, []int{2, 1}}, false},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CanFinish(c.numCourses, c.pres)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t with input %d and %v",
					c.exp, ret, c.numCourses, c.pres)
			}
		})
	}
}

func TestCanFinish2(t *testing.T) {
	st := []struct {
		name       string
		numCourses int
		pres       [][]int
		exp        bool
	}{
		{"zero courses", 0, nil, true},
		{"signle course", 1, nil, true},
		{"no pres", 4, nil, true},
		{"test1", 2, [][]int{[]int{1, 0}}, true},
		{"test4", 2, [][]int{[]int{0, 1}}, true},
		{"test2", 2, [][]int{[]int{1, 0}, []int{0, 1}}, false},
		{"test3", 3, [][]int{[]int{1, 2}, []int{2, 1}}, false},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CanFinish2(c.numCourses, c.pres)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t with input %d and %v",
					c.exp, ret, c.numCourses, c.pres)
			}
		})
	}
}
