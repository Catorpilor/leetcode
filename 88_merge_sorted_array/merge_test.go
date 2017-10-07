package merge

import "testing"
import "reflect"

func TestMerge(t *testing.T){
	st := []struct{
		name string
		nums1, nums2 []int
		m,n int
		exp []int
	}{
		{"test1",[]int{1,2,3,0,0,0,0},[]int{3,4,5,6},3,4,[]int{1,2,3,3,4,5,6}},
		{"test2",[]int{4,5,6,0,0,0},[]int{1,2,3},3,3,[]int{1,2,3,4,5,6}},
	}

	for _,c := range st{
		t.Run(c.name, func(t *testing.T){
			ret := Merge(c.nums1,c.nums2,c.m,c.n)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v and %v",
					c.exp, ret, c.nums1, c.nums2)
			}
		})
	}
}