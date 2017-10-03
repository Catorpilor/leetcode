package add

import "testing"

func TestAddStrings(t *testing.T) {
	st := []struct {
		name, nums1, nums2, exp string
	}{
		{"empty strings", "", "", ""},
		{"one is empty", "123", "", "123"},
		{"one is zero", "9133", "0", "9133"},
		{"non empty without carrier", "123", "456", "579"},
		{"non empty with carrier", "14", "96", "110"},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := AddStrings(c.nums1, c.nums2)
			if ret != c.exp {
				t.Fatalf("expected %s but got %s, with inputs %s and %s",
					c.exp, ret, c.nums1, c.nums2)
			}
		})
	}
}
