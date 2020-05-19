package stock

import "testing"

func constructWithTest(t *testing.T) *StockSpanner {
    ss := Construct()
    t.Cleanup(func() {
        ss = nil
    })
    return ss
}

func TestNext(t *testing.T) {
    ss := constructWithTest(t)
    st := []struct {
        name  string
        price int
        exp   int
    }{
        {"add 100", 100, 1},
        {"add 80", 80, 1},
        {"add 60", 60, 1},
        {"add 70", 70, 2},
        {"add 60", 60, 1},
        {"add 75", 75, 4},
        {"add 85", 85, 6},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := ss.Next(tt.price)
            if out != tt.exp {
                t.Fatalf("after add: %d wanted %d but got %d", tt.price, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
