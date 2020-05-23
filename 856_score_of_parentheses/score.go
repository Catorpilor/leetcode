package score

import "github.com/catorpilor/leetcode/utils"

func scoreOfParenthese(s string) int {
	return useStack(s)
}

// useStack time complexity O(N),space complexity O(N)
func useStack(s string) int {
	n := len(s)
	st := make([]int, 0, n)
	// cur represent the cur score
	var cur int
	for _, c := range s {
		// if we encouter a '(' put current cur to the stack and reset it to 0
		// this is like the AB mode, when we come to B put A's score in the stack
		if c == '(' {
			st = append(st, cur)
			cur = 0
		} else {

			// pop from stack
			nst := len(st)
			top := st[nst-1]
			st = st[:nst-1]
			// top means the scofe of A
			// if cur == 0, which means () we took 1 point and update cur
			// otherwise its (B) case, we took 2*B's score
			cur = top + utils.Max(cur*2, 1)
		}
	}
	return cur
}
