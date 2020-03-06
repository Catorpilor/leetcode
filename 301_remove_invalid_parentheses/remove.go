package remove

func RemoveInvalidP(s string) []string {
	var ret []string
	if len(s) < 2 {
		return ret
	}
	par := [2]byte{'(', ')'}
	helper(s, &ret, 0, 0, par)
	return ret
}

func helper(s string, res *[]string, lasti, lastj int, par [2]byte) {
	for counter, i := 0, lasti; i < len(s); i++ {
		if s[i] == par[0] {
			counter++
		}
		if s[i] == par[1] {
			counter--
		}
		if counter >= 0 {
			continue
		}
		// we have more )s which means counter < 0
		for j := lastj; j <= i; j++ {
			if s[j] == par[1] && (j == lastj || s[j-1] != par[1]) {
				// find 1st ) in the concecutive )s
				//This is for duplication when dealing with case like '())'
				//noted that you are trying out every possible remove of excessive par[1]
				//so in the above case, both ")" will be checked, but the later one shoudn't
				//matter.
				helper(s[:j]+s[j+1:], res, i, j, par)
			}
		}
		//When you are at this point, stack must < 0, so you have no reason to
		//Check with the reverse version. So you simply return
		return
	}
	//When you are at this point, you only know stack must >= 0.
	//And you only finished left to right checked, with deleting the invalid
	//parenthesis from left "()())()" ==> you only delete s[1] and checked (())()
	//you haven't find the case of delete s[4] instead.
	//So at this point, you would need to do the reverse version for the
	//right to left check!
	rvs := reverse(s)
	if par[0] == '(' {
		// reverse s to do another search
		par[0], par[1] = par[1], par[0]
		helper(rvs, res, 0, 0, par)
	} else {
		*res = append(*res, rvs)
	}
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func useBackTracking(s string) []string {
	var left, right int
	for i := range s {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			if left == 0 {
				right++
			} else if left > 0 {
				left--
			}
		}
	}
	var ans []string
	seen := make(map[string]bool)
	backtrack(&ans, s, seen, 0, 0, 0, left, right, []byte{})
	return ans
}

// backtrack worst case "(((((((" time complexity O(2^N)
func backtrack(ans *[]string, s string, seen map[string]bool, idx, leftP, rightP, leftRem, rightRem int, expr []byte) {
	if idx == len(s) {
		if leftRem == 0 && rightRem == 0 && !seen[string(expr)] {
			tmp := make([]byte, len(expr))
			copy(tmp, expr)
			st := string(tmp)
			seen[st] = true
			*ans = append((*ans), st)
		}
	} else {
		ch := s[idx]
		curLen := len(expr)
		if (ch == '(' && leftRem > 0) || (ch == ')' && rightRem > 0) {
			// remove case
			if ch == '(' {
				backtrack(ans, s, seen, idx+1, leftP, rightP, leftRem-1, rightRem, expr)
			} else {
				backtrack(ans, s, seen, idx+1, leftP, rightP, leftRem, rightRem-1, expr)
			}
		}
		// append
		expr = append(expr, ch)
		if ch != '(' && ch != ')' {
			backtrack(ans, s, seen, idx+1, leftP, rightP, leftRem, rightRem, expr)
		} else if ch == '(' {
			backtrack(ans, s, seen, idx+1, leftP+1, rightP, leftRem, rightRem, expr)
		} else if rightP < leftP {
			backtrack(ans, s, seen, idx+1, leftP, rightP+1, leftRem, rightRem, expr)
		}
		// backtrack
		expr = expr[:curLen]
	}
}
