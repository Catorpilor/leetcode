package uf

import "testing"

var qf *QuickFind

func setupQF(n int) func() {
    if qf == nil {
        qf = NewQuickFind(n)
    }
    return func() {
        qf = nil
    }
}

func TestQFFind(t *testing.T) {
    f := setupQF(10)
    defer f()
    qf.Union(0, 1)
    qf.Union(2, 3)
    qf.Union(1, 9)
    if qf.Find(0, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    if qf.Find(2, 9) != false {
        t.Fatal("uniond 0,1,9 should got false but got true")
    }
    if qf.Find(1, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    t.Log("pass")
}

func TestQFUnion(t *testing.T) {
    f := setupQF(10)
    defer f()
    if qf.Find(0, 1) != false {
        t.Fatal("uniond 0,1,9 should got false but got true")
    }
    qf.Union(0, 1)
    if qf.Find(0, 1) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    qf.Union(2, 3)
    qf.Union(1, 9)
    if qf.Find(0, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    if qf.Find(2, 9) != false {
        t.Fatal("uniond 0,1,9 should got false but got true")
    }
    if qf.Find(1, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    t.Log("pass")
}
