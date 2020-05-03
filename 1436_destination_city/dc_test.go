package dc

import "testing"

func TestDestCity(t *testing.T) {
    st := []struct {
        name  string
        paths [][]string
        exp   string
    }{
        {"empty paths", nil, ""},
        {"single path", [][]string{[]string{"A", "B"}}, "B"},
        {"testcase1", [][]string{[]string{"London", "New York"}, []string{"New York", "Lima"}, []string{"Lima", "Sao Paulo"}}, "Sao Paulo"},
        {"testcase2", [][]string{[]string{"B", "C"}, []string{"D", "B"}, []string{"C", "A"}}, "A"},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := destCity(tt.paths)
            if out != tt.exp {
                t.Fatalf("with input paths: %v wanted %s but got %s", tt.paths, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
