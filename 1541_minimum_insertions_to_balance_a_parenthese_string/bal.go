package bal

func minInserts(s string) int {
	return useStack(s)
}

// useStack time complexity O(N), space complexity O(N)
func useStack(s string) int {
	n := len(s)
	st := make([]int, 0, n)
	var ans int
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			st = append(st, i)
			continue
		}
		nst := len(st)
		if i < n-1 && s[i+1] == ')' {
			if nst > 0 {
				st = st[:nst-1]
			} else {
				// add one left parenthese
				ans++
			}
			// update i skip 1
			i++
		} else {
			if nst > 0 {
				st = st[:nst-1]
				// add one right parenthese
				ans++
			} else {
				// add one left parenthese and right parentheses
				ans += 2
			}
		}
	}

	return ans + 2*len(st)
}
