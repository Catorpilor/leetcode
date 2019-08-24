package att

import (
    "fmt"
    "strings"
)

func CheckRecord(n int) int {
    var cnt int
    if n > 0 {
        // bruteForce(&cnt, n)
        // cnt = RecursiveFormula(n)
        cnt = DynamicProg(n)
    }
    return cnt
}

// time complexity O(3^n)
// space O(n^n)
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

const (
    M = 1000000007
)

// time complexity O(2^n)
// space O(n)
func RecursiveFormula(n int) int {
    f := make([]int, n+1)
    f[0] = 1
    for i := 1; i <= n; i++ {
        f[i] = recur(i)
    }
    cnt := f[n]
    for i := 1; i <= n; i++ {
        cnt += (recur(i-1) * recur(n-i)) % M
    }
    return cnt % M
}

func recur(n int) int {
    switch n {
    case 0:
        return 1
    case 1:
        return 2
    case 2:
        return 4
    case 3:
        return 7 // 2^3 - 1 (LLL)
    default:
        return (2*recur(n-1) - recur(n-4)) % M
    }
}

func DynamicProg(n int) int {
    capN := n + 1
    if n <= 5 {
        capN = 6
    }
    dp := make([]int, capN)
    dp[0] = 1
    dp[1] = 2
    dp[2] = 4
    dp[3] = 7
    for i := 4; i <= n; i++ {
        dp[i] = (2*dp[i-1])%M - (dp[i-4])%M
    }
    sum := dp[n]
    for i := 1; i <= n; i++ {
        sum += (dp[i-1] * dp[n-i]) % M
    }
    return sum % M

}
