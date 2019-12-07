package ltp

import "testing"

func TestNumOfPossibilities(t *testing.T) {
    st := []struct {
        name  string
        tiles string
        exp   int
    }{
        {"mepty", "", 0},
        {"single", "A", 1},
        {"identical", "AAA", 3},
        {"testcase1", "AAB", 8},
        {"testcase2", "AAABBC", 188},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := numOfPossibilities(tt.tiles)
            if out != tt.exp {
                t.Fatalf("with tiles: %s wanted %d but got %d", tt.tiles, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestNumOfPosBt(t *testing.T) {
    st := []struct {
        name  string
        tiles string
        exp   int
    }{
        {"mepty", "", 0},
        {"single", "A", 1},
        {"identical", "AAA", 3},
        {"testcase1", "AAB", 8},
        {"testcase2", "AAABBC", 188},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := numOfPossibilities2(tt.tiles)
            if out != tt.exp {
                t.Fatalf("with tiles: %s wanted %d but got %d", tt.tiles, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
