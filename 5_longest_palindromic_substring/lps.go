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
	dp[0] = 1
	// prevcenter, leftbound, rightbound
	pc, lb, rb, max := 0, 0, 0, 1
	for i := 1; i < cap; i++ {
		cnt := maxPalinLength(exs, i, i)
	}

}

func maxPalinLength(b []byte, l, r int) int {
	var count int
	for l >= 0 && r < len(b) && b[l] == b[r] {
		count, l, r = count+1, l-1, r+1
	}
	return count
}
