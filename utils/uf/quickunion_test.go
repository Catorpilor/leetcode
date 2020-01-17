package uf

import "testing"

var qu *QuickUnion

func setupQU(n int) func() {
    if qu == nil {
        qu = NewQuickUnion(n)
    }
    return func() {
        qu = nil
    }
}

func TestQUFind(t *testing.T) {
    f := setupQU(10)
    defer f()
    qu.Union(0, 1)
    qu.Union(2, 3)
    qu.Union(1, 9)
    if qu.Find(0, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    if qu.Find(2, 9) != false {
        t.Fatal("uniond 0,1,9 should got false but got true")
    }
    if qu.Find(1, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    t.Log("pass")
}

func TestQUUnion(t *testing.T) {
    f := setupQU(10)
    defer f()
    if qu.Find(0, 1) != false {
        t.Fatal("uniond 0,1,9 should got false but got true")
    }
    qu.Union(0, 1)
    if qu.Find(0, 1) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    qu.Union(2, 3)
    qu.Union(1, 9)
    if qu.Find(0, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    if qu.Find(2, 9) != false {
        t.Fatal("uniond 0,1,9 should got false but got true")
    }
    if qu.Find(1, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    t.Log("pass")
}
