package validparenthesis

import "testing"

func TestIsValidParenthesis(t *testing.T) {
    st := []struct {
        name string
        s    string
        exp  bool
    }{
        {"empty string", "", true},
        {"just a *", "*", true},
        {"testcase1", ")(", false},
        {"testcase2", "())()", false},
        {"testcase3", "(*)*(((*", false},
        {"testcase4", "(()((*)*))", true},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := isvalidParenthesis(tt.s)
            if out != tt.exp {
                t.Fatalf("with input s:%s wanted %t but got %t", tt.s, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestIsValidParenthesisUseDfs(t *testing.T) {
    st := []struct {
        name string
        s    string
        exp  bool
    }{
        {"empty string", "", true},
        {"just a *", "*", true},
        {"testcase1", ")(", false},
        {"testcase2", "())()", false},
        {"testcase3", "(*)*(((*", false},
        {"testcase4", "(()((*)*))", true},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := useDfsWithMemorization(tt.s)
            if out != tt.exp {
                t.Fatalf("with input s:%s wanted %t but got %t", tt.s, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
