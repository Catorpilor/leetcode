package lcp

import (
    "fmt"
    "strings"
)

func letterCombinations(digits string) []string {
    phone := map[byte]string{
        '1': "",
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "prqs",
        '8': "tuv",
        '9': "wxyz",
        '0': "",
    }
    var res []string
    if strings.Index(digits, "1") != -1 || strings.Index(digits, "0") != -1 || len(digits) == 0 {
        return res
    }
    permute(&res, phone, "", digits, 0)
    return res
}

// pos = 0 digits = "23"
// bid = "a"
// pos = 1
// bid = "ad"
// append res = [ad]

// permute process the backtracking process
// kmap phone keyboard mapping
// bid current key combination
// pos current processing digits[pos]
func permute(res *[]string, kmap map[byte]string, bid, digits string, pos int) {
    if len(bid) == len(digits) {
        fmt.Printf("bid: %s\n", bid)
        *res = append(*res, bid)
        return
    }
    for i := pos; i < len(digits); i++ {
        vals := kmap[digits[i]]
        for j := 0; j < len(vals); j++ {
            n := len(bid)
            bid += string(vals[j])
            permute(res, kmap, bid, digits, i+1)
            bid = bid[:n]
            fmt.Printf("after permute bid[:n]: %s, pos: %d, i: %d\n", bid, pos, i)
        }
    }
}
