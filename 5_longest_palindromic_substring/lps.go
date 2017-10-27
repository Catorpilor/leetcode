package lps

// Lps returns the longest papindromic substring in s
func Lps(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	start, maxLen := 0, 1
	// for lenght 1
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	// for length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			start, maxLen = i, 2
			dp[i][i+1] = true
		}
	}
	// for length >= 3
	for cl := 3; cl <= n; cl++ {
		for i := 0; i < n-cl+1; i++ {
			j := i + cl - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				maxLen = cl
			}
		}
	}
	return s[start : start+maxLen]
}

// Lps2 returns the longest palindromic substring in s
func Lps2(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	cap, extended, exs := n, false, []byte(s)
	if n%2 == 0 {
		cap = 2*n + 1
		extended = true
		exs = nil
		exs = make([]byte, cap)
		for i := range exs {
			if i%2 == 0 {
				exs[i] = '$'
			} else {
				exs[i] = s[i/2]
			}
		}
	}
	dp := make([]int, cap)
	for i := range dp {
		dp[i] = 1
	}
	// prevcenter, leftbound, rightbound
	pc, lb, rb, max, nc, fc := 0, 0, 0, 1, 1, 0
	// i stands for current cent
	for nc < cap {
		cnt := maxPalinLength(exs, nc-dp[nc]/2-1, nc+dp[nc]/2+1)
		dp[nc] += cnt
		if dp[nc] > max {
			max = dp[nc]
			fc = nc
		}
		lb, rb = nc-dp[nc]/2, nc+dp[nc]/2
		if rb-lb <= 3 {
			nc += 1
		} else {
			// we have to find next center
			if rb == cap-1 {
				break
			}
			pc = nc
			cmax := dp[nc-1]
			for j, k := nc+1, nc-1; j <= rb && k >= lb; j, k = j+1, k-1 {
				dp[j] = dp[k]
				if k-dp[k]/2 < lb {
					if k+dp[k]/2 > nc {
						dp[j] = rb - nc + 1
					} else {
						dp[j] = 1
					}
				}
				if dp[j] > cmax && j+dp[j]/2 >= rb && k-dp[k]/2 >= lb {
					nc, cmax = j, dp[k]
				}
			}
			// if not found
			if pc == nc {
				nc = rb
			}
		}
	}
	start, end := fc-max/2, fc+max/2+1
	if extended {
		max /= 2
		start, end = fc-max, fc
		if fc%2 != 0 && start > 1 {
			start, end = start-1, end-1
		}
	}
	if end >= n {
		return s[start:]
	}
	return s[start:end]
}

func maxPalinLength(b []byte, l, r int) int {
	count := 0
	for l >= 0 && r < len(b) && b[l] == b[r] {
		count, l, r = count+2, l-1, r+1
	}
	return count
}
