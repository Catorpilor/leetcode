package desc

import "testing"

func TestDesc(t *testing.T){
	st := []struct{
		name string
		nums []int
		exp bool
	}{
		{"nil slice",nil, true},
		{"single element",[]int{1},true},
		{"descring",[]int{4,2,1},false},
		{"can change",[]int{4,2,3},true},
		{"3423",[]int{3,4,2,3},false},
	}

	for _,c := range st{
		t.Run(c.name, func(t *testing.T){
			ret := descWithoutMidification(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}