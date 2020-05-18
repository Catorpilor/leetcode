package words

import (
    "sort"
    "strings"
    "unicode"
)

func rearrangeWords(text string) string {
    return useBruteForce(text)
}

// useBruteForce time complexity O(N), space complexity O(10000*MaxLen)
func useBruteForce(text string) string {
    arr := strings.Fields(strings.ToLower(text))
    set := make([][]string, 10_000)
    for i := range arr {
        set[len(arr[i])] = append(set[len(arr[i])], arr[i])
    }
    ans := make([]string, 0, len(arr))
    for i := range set {
        if len(set[i]) > 0 {
            ans = append(ans, set[i]...)
        }
    }
    // sort.Slice(arr, func(i, j int) bool {
    // return len(arr[i]) < len(arr[j])
    // })
    for i := range ans {
        bc := []byte(ans[i])
        if i == 0 {
            if bc[0] >= 'a' && bc[0] <= 'z' {
                up := byte(unicode.ToUpper(rune(bc[0])))
                bc[0] = up
                //arr[i] = string(bc)
            }
        }
        ans[i] = string(bc)
    }
    return strings.Join(ans, " ")
}

type stringSlice []string

func (ss stringSlice) Len() int {
    return len(ss)
}

func (ss stringSlice) Less(i, j int) bool {
    return len(ss[i]) < len(ss[j])
}

func (ss stringSlice) Swap(i, j int) {
    ss[i], ss[j] = ss[j], ss[i]
}

// useStableSort time complexity O(N*lgN), space compelxity O(1)
func useStableSort(text string) string {
    arr := strings.Fields(strings.ToLower(text))
    sort.Stable(stringSlice(arr))
    bc := []byte(arr[0])
    up := byte(unicode.ToUpper(rune(bc[0])))
    bc[0] = up
    arr[0] = string(bc)
    return strings.Join(arr, " ")
}
