package max

import "testing"

func TestLargestPalindrome(t *testing.T){
	st := []struct{
		name string
		n,exp int
	}{
		{"n=1",1,9},
		{"n=2",2,987},
		{"n=3",3,123},
		{"n=4",4,597},
		{"n=5",5,677},
		{"n=6",6,1218},
		{"n=7",7,877},
		{"n=8",8,475},
		{"n=9",9,1226},
	}

	for _,c := range st {
		t.Run(c.name, func(t *testing.T){
			ret := LargestPalindrome(c.n)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %d",
					c.exp, ret, c.n)
			}
		})
	}

}
