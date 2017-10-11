package stro

import "testing"

func TestIsStro(t *testing.T) {
	st := []struct {
		name, num string
		exp       bool
	}{
		{"1", "1", true},
		{"0", "0", true},
		{"12", "12", false},
		{"69", "69", true},
		{"818", "818", true},
		{"8417", "8417", false},
	}
	for _, c := range st {
		ret := IsStro(c.num)
		if ret != c.exp {
			t.Fatalf("expected %t but got %t with input %s",
				c.exp, ret, c.num)
		}
	}
}
