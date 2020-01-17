package uf

import "testing"

var wqu *WeightedQuickUnion

func setupWQU(n int) func() {
    if wqu == nil {
        wqu = NewWeightedQuickUnion(n)
    }
    return func() {
        wqu = nil
    }
}

func TestWQUFind(t *testing.T) {
    f := setupWQU(10)
    defer f()
    wqu.Union(0, 1)
    wqu.Union(2, 3)
    wqu.Union(1, 9)
    if wqu.Find(0, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    if wqu.Find(2, 9) != false {
        t.Fatal("uniond 0,1,9 should got false but got true")
    }
    if wqu.Find(1, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    t.Log("pass")
}

func TestWQUUnion(t *testing.T) {
    f := setupWQU(10)
    defer f()
    if wqu.Find(0, 1) != false {
        t.Fatal("uniond 0,1,9 should got false but got true")
    }
    wqu.Union(0, 1)
    if wqu.Find(0, 1) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    wqu.Union(2, 3)
    wqu.Union(1, 9)
    if wqu.Find(0, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    if wqu.Find(2, 9) != false {
        t.Fatal("uniond 0,1,9 should got false but got true")
    }
    if wqu.Find(1, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    t.Log("pass")
}
