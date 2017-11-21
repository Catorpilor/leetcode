package meeting

import "testing"

func TestMinRooms(t *testing.T) {
	st := []struct {
		name      string
		intervals []interval
		exp       int
	}{
		{"empty intervals", []interval{}, 0},
		{"single", []interval{interval{1, 2}}, 1},
		{"no overlapping", []interval{interval{1, 2}, interval{3, 4}}, 1},
		{"with overlapping", []interval{interval{0, 20}, interval{5, 10}, interval{6, 8}}, 3},
		{"failed 1", []interval{interval{1, 5}, interval{8, 9}, interval{8, 9}}, 2},
		{"failed 2", []interval{interval{4, 19}, interval{6, 13}, interval{11, 20}, interval{13, 17}}, 3},
		{"failed 3", []interval{interval{2, 36}, interval{3, 4}, interval{13, 34}, interval{16, 20}, interval{39, 40}}, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MinRooms(c.intervals)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.intervals)
			}
		})
	}
}

func TestMinRooms2(t *testing.T) {
	st := []struct {
		name      string
		intervals []interval
		exp       int
	}{
		{"empty intervals", []interval{}, 0},
		{"single", []interval{interval{1, 2}}, 1},
		{"no overlapping", []interval{interval{1, 2}, interval{3, 4}}, 1},
		{"with overlapping", []interval{interval{0, 20}, interval{5, 10}, interval{6, 8}}, 3},
		{"failed 1", []interval{interval{1, 5}, interval{8, 9}, interval{8, 9}}, 2},
		{"failed 2", []interval{interval{4, 19}, interval{6, 13}, interval{11, 20}, interval{13, 17}}, 3},
		{"failed 3", []interval{interval{2, 36}, interval{3, 4}, interval{13, 34}, interval{16, 20}, interval{39, 40}}, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MinRooms2(c.intervals)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.intervals)
			}
		})
	}
}

func TestMinRooms3(t *testing.T) {
	st := []struct {
		name      string
		intervals []interval
		exp       int
	}{
		{"empty intervals", []interval{}, 0},
		{"single", []interval{interval{1, 2}}, 1},
		{"no overlapping", []interval{interval{1, 2}, interval{3, 4}}, 1},
		{"with overlapping", []interval{interval{0, 20}, interval{5, 10}, interval{6, 8}}, 3},
		{"failed 1", []interval{interval{1, 5}, interval{8, 9}, interval{8, 9}}, 2},
		{"failed 2", []interval{interval{4, 19}, interval{6, 13}, interval{11, 20}, interval{13, 17}}, 3},
		{"failed 3", []interval{interval{2, 36}, interval{3, 4}, interval{13, 34}, interval{16, 20}, interval{39, 40}}, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MinRooms3(c.intervals)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.intervals)
			}
		})
	}
}

func TestMinRooms4(t *testing.T) {
	st := []struct {
		name      string
		intervals []interval
		exp       int
	}{
		{"empty intervals", []interval{}, 0},
		{"single", []interval{interval{1, 2}}, 1},
		{"no overlapping", []interval{interval{1, 2}, interval{3, 4}}, 1},
		{"with overlapping", []interval{interval{0, 20}, interval{5, 10}, interval{6, 8}}, 3},
		{"failed 1", []interval{interval{1, 5}, interval{8, 9}, interval{8, 9}}, 2},
		{"failed 2", []interval{interval{4, 19}, interval{6, 13}, interval{11, 20}, interval{13, 17}}, 3},
		{"failed 3", []interval{interval{2, 36}, interval{3, 4}, interval{13, 34}, interval{16, 20}, interval{39, 40}}, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MinRooms4(c.intervals)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.intervals)
			}
		})
	}
}
