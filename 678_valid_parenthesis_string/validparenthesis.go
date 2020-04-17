package validparenthesis

import "github.com/catorpilor/leetcode/utils"

func isvalidParenthesis(s string) bool {
    return useGreedy(s)
}

// useGreedy time complexity O(N), space complexity O(1)
func useGreedy(s string) bool {
    // in the end we only have to count how many left parenthesis left.
    // for example when traverse string "(***)" we calculate the possbile value of
    // "("
    // "(" --------> [1]
    // "(*" -------> [0,1,2], 0 means "*" = ")", 1 means "*"="", 2 means "*"="("
    // "(**"-------> [0,1,2,3]
    // "(***"------> [0,1,2,3,4]
    // "(***)" ----> [0,1,2,3]
    // we can use two pointers to simplify the above range [low, hi]
    var low, hi int
    for i := range s {
        if s[i] == '(' {
            low++
        } else {
            low--
        }
        if s[i] != ')' {
            hi++
        } else {
            hi--
        }
        if hi < 0 {
            break
        }
        low = utils.Max(low, 0)
    }
    return low == 0
}

// useDfsWithMemorization time complexity O(N)?, space compleixty O(N^2)
func useDfsWithMemorization(s string) bool {
    n := len(s)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n)
    }
    return helper(&memo, s, 0, 0)
}

func helper(memo *[][]int, s string, idx, count int) bool {
    if count < 0 {
        return false
    }
    if idx >= len(s) {
        if count == 0 {
            return true
        }
        return false
    }
    if (*memo)[idx][count] != 0 {
        if (*memo)[idx][count] == 1 {
            return true
        }
        return false
    }
    var ret bool
    if s[idx] == '(' {
        ret = ret || helper(memo, s, idx+1, count+1)
    } else if s[idx] == ')' {
        ret = ret || helper(memo, s, idx+1, count-1)
    } else {
        ret = ret || helper(memo, s, idx+1, count) || helper(memo, s, idx+1, count+1) || helper(memo, s, idx+1, count-1)
    }
    if ret {
        (*memo)[idx][count] = 1
    } else {
        (*memo)[idx][count] = -1
    }
    return ret
}
