package att

import (
    "fmt"
    "strings"
)

func CheckRecord(n int) int {
    var cnt int
    if n > 0 {
        bruteForce(&cnt, n)
    }
    return cnt
}

func bruteForce(count *int, n int) {
    gen("", count, n)
    return
}

func gen(s string, count *int, n int) {
    if n == 0 && isValid(s) {
        fmt.Printf("current valid s: %s\n", s)
        *count++
    } else if n > 0 {
        fmt.Printf("prev S: %s, count: %d, n : %d\n", s, *count, n)
        gen(s+"A", count, n-1)
        gen(s+"L", count, n-1)
        gen(s+"P", count, n-1)
    }
}

func isValid(s string) bool {
    var cnt int
    for i := 0; i < len(s) && cnt < 2; i++ {
        if s[i] == 'A' {
            cnt++
        }
    }
    return len(s) > 0 && cnt < 2 && strings.Index(s, "LLL") < 0
}
