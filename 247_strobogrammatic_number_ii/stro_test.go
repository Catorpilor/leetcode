package stro

import "testing"

func TestFindStro(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  []string
	}{
		{"n eq 0", 0, []string{""}},
		{"1", 1, []string{"0", "1", "8"}},
		{"2", 2, []string{"11", "69", "88", "96"}},
	}
	for _, c := range st {
		ret := FindStro(c.n)
		if len(ret) != len(c.exp) {
			t.Fatalf("expected %v, but got %v, with input %d",
				c.exp, ret, c.n)
		}
	}
}

func TestFindStro2(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  []string
	}{
		{"n eq 0", 0, []string{""}},
		{"1", 1, []string{"0", "1", "8"}},
		{"2", 2, []string{"11", "69", "88", "96"}},
	}
	for _, c := range st {
		ret := FindStro2(c.n)
		if len(ret) != len(c.exp) {
			t.Fatalf("expected %v, but got %v, with input %d",
				c.exp, ret, c.n)
		}
	}
}
