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
