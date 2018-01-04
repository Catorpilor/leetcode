package matrix

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	st := []struct {
		name   string
		matrix [][]int
		exp    [][]int
	}{
		{"empty", [][]int{}, [][]int{}},
		{"single0", [][]int{[]int{0}}, [][]int{[]int{0}}},
		{"test1", [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}},
		{"test2", [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 2, 1}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := UpdateMatrix(c.matrix)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.matrix)
			}
		})
	}
}

func TestUpdateMatrix2(t *testing.T) {
	st := []struct {
		name   string
		matrix [][]int
		exp    [][]int
	}{
		{"empty", [][]int{}, [][]int{}},
		{"single0", [][]int{[]int{0}}, [][]int{[]int{0}}},
		{"test1", [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}},
		{"test2", [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 2, 1}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := UpdateMatrix2(c.matrix)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.matrix)
			}
		})
	}
}

func TestUpdateMatrix3(t *testing.T) {
	st := []struct {
		name   string
		matrix [][]int
		exp    [][]int
	}{
		{"empty", [][]int{}, [][]int{}},
		{"single0", [][]int{[]int{0}}, [][]int{[]int{0}}},
		{"test1", [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}},
		{"test2", [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 2, 1}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := UpdateMatrix3(c.matrix)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.matrix)
			}
		})
	}
}
