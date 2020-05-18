package words

import "testing"

func TestRearrangeWords(t *testing.T) {
    st := []struct {
        name string
        text string
        exp  string
    }{
        {"just a single Word", "As", "As"},
        {"testcase1", "Leetcode is cool", "Is cool leetcode"},
        {"testcase2", "Keep calm and code on", "On and keep calm code"},
        {"testcase3", "To be or not to be", "To be or to be not"},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := rearrangeWords(tt.text)
            if out != tt.exp {
                t.Fatalf("with input text:%s wanted %s but got %s", tt.text, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
