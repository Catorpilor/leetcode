package cal

import "testing"

var calender = NewCalender()

func TestBook(t *testing.T) {
	st := []struct {
		name       string
		start, end int
		exp        bool
	}{
		{"book1", 47, 50, true},
		{"book2", 33, 41, true},
		{"book3", 39, 45, false},
		{"book4", 33, 42, false},
		{"book4", 25, 32, true},
		{"book4", 26, 35, false},
		{"book4", 19, 25, true},
		{"book4", 3, 8, true},
		{"book4", 8, 13, true},
		{"book4", 18, 27, false},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			// calender.Status()
			ret := calender.Book(c.start, c.end)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t with input %d and %d",
					c.exp, ret, c.start, c.end)
			}
		})
	}
}

func TestBook2(t *testing.T) {
	st := []struct {
		name       string
		start, end int
		exp        bool
	}{
		{"book1", 47, 50, true},
		{"book2", 33, 41, true},
		{"book3", 39, 45, false},
		{"book4", 33, 42, false},
		{"book4", 25, 32, true},
		{"book4", 26, 35, false},
		{"book4", 19, 25, true},
		{"book4", 3, 8, true},
		{"book4", 8, 13, true},
		{"book4", 18, 27, false},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			// calender.Status()
			ret := calender.Book(c.start, c.end)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t with input %d and %d",
					c.exp, ret, c.start, c.end)
			}
		})
	}
}
