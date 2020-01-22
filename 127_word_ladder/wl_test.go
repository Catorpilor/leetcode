package wl

import "testing"

func TestLadderLength(t *testing.T) {
    st := []struct {
        name       string
        begin, end string
        wordList   []string
        exp        int
    }{
        {"empty wordList", "abc", "cde", []string{}, 0},
        {"testcase1", "hit", "cog", []string{"hot", "dot", "dog", "log", "cog"}, 5},
        {"testcase2", "hit", "cog", []string{"hot", "dot", "dog", "log"}, 0},
        {"failed1", "a", "c", []string{"a", "b", "c"}, 2},
        {"failed2", "hot", "dot", []string{"hot", "dot", "dog"}, 2},
        {"failed3", "leet", "code", []string{"lest", "leet", "lose", "code", "lode", "robe", "lost"}, 6},
        {"failed4", "hog", "cog", []string{"cog"}, 2},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := ladderLength(tt.begin, tt.end, tt.wordList)
            if out != tt.exp {
                t.Fatalf("with worldList: %v and begin:%s and end: %s wanted %d but got %d",
                    tt.wordList, tt.begin, tt.end, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestDiffOne(t *testing.T) {
    st := []struct {
        name string
        a, b string
        exp  bool
    }{
        {"testcase1", "abc", "abd", true},
        {"testcase2", "aac", "aab", true},
        {"testcase3", "abc", "bdc", true},
        {"testcase4", "abc", "bfe", false},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := diffOne(tt.a, tt.b)
            if out != tt.exp {
                t.Fatalf("with a:%s and b:%s wanted %t but got %t", tt.a, tt.b, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
