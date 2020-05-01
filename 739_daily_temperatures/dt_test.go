package dt

import (
    "reflect"
    "testing"
)

func TestDailyTemps(t *testing.T) {
    st := []struct {
        name  string
        temps []int
        exp   []int
    }{
        {"single temp", []int{34}, []int{0}},
        {"all identical", []int{33, 33, 33}, []int{0, 0, 0}},
        {"testcase1", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := dailyTemps(tt.temps)
            if !reflect.DeepEqual(out, tt.exp) {
                t.Fatalf("with input temps:%v wanted %v but got %v", tt.temps, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
