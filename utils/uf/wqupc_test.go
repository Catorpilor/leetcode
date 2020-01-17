package uf

import "testing"

var wqupc *WQUPC

func setupWQUPC(n int) func() {
    if wqupc == nil {
        wqupc = NewWQUPC(n)
    }
    return func() {
        wqupc = nil
    }
}

func TestWQUPCFind(t *testing.T) {
    f := setupWQUPC(10)
    defer f()
    wqupc.Union(0, 1)
    wqupc.Union(2, 3)
    wqupc.Union(1, 9)
    if wqupc.Find(0, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    if wqupc.Find(2, 9) != false {
        t.Fatal("uniond 0,1,9 should got false but got true")
    }
    if wqupc.Find(1, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    t.Log("pass")
}

func TestWQUPCUnion(t *testing.T) {
    f := setupWQUPC(10)
    defer f()
    if wqupc.Find(0, 1) != false {
        t.Fatal("uniond 0,1,9 should got false but got true")
    }
    wqupc.Union(0, 1)
    if wqupc.Find(0, 1) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    wqupc.Union(2, 3)
    wqupc.Union(1, 9)
    if wqupc.Find(0, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    if wqupc.Find(2, 9) != false {
        t.Fatal("uniond 0,1,9 should got false but got true")
    }
    if wqupc.Find(1, 9) != true {
        t.Fatal("uniond 0,1,9 should got true but got false")
    }
    t.Log("pass")
}
