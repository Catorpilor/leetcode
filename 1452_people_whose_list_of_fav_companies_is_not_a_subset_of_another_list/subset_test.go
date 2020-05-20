package subset

import (
    "reflect"
    "testing"
)

func TestNotASubSet(t *testing.T) {
    st := []struct {
        name         string
        favCompanies [][]string
        exp          []int
    }{
        {"testcaes1", [][]string{[]string{"google"}, []string{"leetcode"}, []string{"amazon"}}, []int{0, 1, 2}},
        {"testcase2", [][]string{[]string{"leetcode", "google", "facebook"}, []string{"google", "microsoft"}, []string{"google", "facebook"},
            []string{"google"}, []string{"amazon"}}, []int{0, 1, 4}},
        {"testcaes3", [][]string{[]string{"leetcode", "google", "facebook"}, []string{"leetcode", "amazon"}, []string{"facebook", "google"}}, []int{0, 1}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := peopleIndexes(tt.favCompanies)
            if !reflect.DeepEqual(out, tt.exp) {
                t.Fatalf("with input favCompanies: %v wanted %v but got %v", tt.favCompanies, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestNotASubSetUseUF(t *testing.T) {
    st := []struct {
        name         string
        favCompanies [][]string
        exp          []int
    }{
        {"testcaes1", [][]string{[]string{"google"}, []string{"leetcode"}, []string{"amazon"}}, []int{0, 1, 2}},
        {"testcase2", [][]string{[]string{"leetcode", "google", "facebook"}, []string{"google", "microsoft"}, []string{"google", "facebook"},
            []string{"google"}, []string{"amazon"}}, []int{0, 1, 4}},
        {"testcaes3", [][]string{[]string{"leetcode", "google", "facebook"}, []string{"leetcode", "amazon"}, []string{"facebook", "google"}}, []int{0, 1}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := useUnionFind(tt.favCompanies)
            if !reflect.DeepEqual(out, tt.exp) {
                t.Fatalf("with input favCompanies: %v wanted %v but got %v", tt.favCompanies, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
