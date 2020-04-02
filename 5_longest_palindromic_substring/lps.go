package lps

import "github.com/catorpilor/leetcode/utils"

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
	cap := 2*n + 1
	exs := make([]byte, cap)
	for i := range exs {
		if i%2 == 0 {
			exs[i] = '$'
		} else {
			exs[i] = s[i/2]
		}
	}

	dp := make([]int, cap)
	for i := range dp {
		dp[i] = 1
	}
	// prevcenter, leftbound, rightbound, maxlength of palindromic substring
	// nc stands for current center, fc stands for final palindromic substring center
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
				// if right bound reaches the end of string
				break
			}
			// set prev cent to current center
			pc = nc
			cmax := 0
			for j, k := nc+1, nc-1; j <= rb && k >= lb; j, k = j+1, k-1 {
				dp[j] = dp[k]
				// if k's left bound goes beyond current center's left bound
				// j can not be our next center
				if k-dp[k]/2 < lb {
					if k+dp[k]/2 > nc {
						dp[j] = rb - nc
					} else {
						dp[j] = 1
					}
				}
				if dp[j] > cmax && j+dp[j]/2 >= rb && k-dp[k]/2 >= lb {
					nc, cmax = j, dp[k]
				}
			}
			// if not found
			// nextcenter should be the right bound of current palindromic substring
			// or rb+1
			if pc == nc {
				nc = rb
			}
		}
	}
	start, end := fc-max/2, fc+max/2+1
	ret := make([]byte, 0, max/2)
	for i := start; i < end; i++ {
		if exs[i] != '$' {
			ret = append(ret, exs[i])
		}
	}
	return string(ret)
}

func maxPalinLength(b []byte, l, r int) int {
	count := 0
	for l >= 0 && r < len(b) && b[l] == b[r] {
		count, l, r = count+2, l-1, r+1
	}
	return count
}

func isPalindrome(s string) bool {
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func isGood(x int, s string) int {
	for l := 0; l+x <= len(s); l++ {
		if isPalindrome(s[l : l+x]) {
			return l
		}
	}
	return -1
}

// useBinarySearch time compleixty O(N^2 * lg(N)), space complexity O(1)
func useBinarySearch(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	var startPos, maxLen int
	for parity := 0; parity < 2; parity++ {
		low, high := 1, n
		if low%2 != parity {
			low++
		}
		if high%2 != parity {
			high--
		}
		for low <= high {
			mid := low + (high-low)/2
			if mid%2 != parity {
				mid++
			}
			if mid > high {
				break
			}
			tmp := isGood(mid, s)
			if tmp != -1 {
				// found one
				if mid > maxLen {
					maxLen = mid
					startPos = tmp
				}
				low = mid + 2
			} else {
				high = mid - 2
			}
		}
	}
	return s[startPos : startPos+maxLen]
}

// useExpendingFromCenter time complexity O(N^2), space complexity O(1)
func useExpendingFromCenter(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	var startPos, maxLen int
	// odd ones
	for x := 0; x < n; x++ {
		for k := 0; x-k >= 0 && x+k < n; k++ {
			if s[x-k] != s[x+k] {
				break
			}
			tmp := 2*k + 1
			// fmt.Printf("odd ones startPos:%d, maxLen:%d, x:%d, k:%d, left:%d, right:%d\n", startPos, maxLen, x, k, x-k, x+k)
			if tmp > maxLen {
				maxLen = tmp
				startPos = x - k
				// fmt.Printf("odd found: startPos:%d, maxLen:%d\n", startPos, maxLen)
			}
		}
	}
	// even palindromes
	// for example
	// s = "a b b a"
	//       ^ ^ ^
	// x=0   | | |
	// x=1     | |
	// x=2       |
	for x := 0; x < n-1; x++ {
		for k := 1; x+1-k >= 0 && x+k < n; k++ {
			if s[x+1-k] != s[x+k] {
				break
			}
			// fmt.Printf("even ones startPos:%d, maxLen:%d, x:%d, k:%d, left:%d, right:%d\n", startPos, maxLen, x, k, x+1-k, x+k)
			tmp := 2 * k
			if tmp > maxLen {
				maxLen = tmp
				startPos = x + 1 - k
				// fmt.Printf("even found: startPos:%d, maxLen:%d\n", startPos, maxLen)
			}
		}
	}
	return s[startPos : startPos+maxLen]
}

func preProcess(s string) []byte {
	n := len(s)
	sb := make([]byte, 2*n+3)
	sb[0] = '$'
	sb[2*n+2] = '@'
	for i := range s {
		sb[2*i+1] = '#'
		sb[2*i+2] = s[i]
	}
	sb[2*n+1] = '#'
	return sb
}

// useMancher time complexity O(N), space complexity O(N)
func useMancher(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	sb := preProcess(s)
	nsb := len(sb)
	lpc := make([]int, nsb)
	var center, right int
	for i := 1; i < nsb-1; i++ {
		mirror := 2*center - i
		if right > i {
			lpc[i] = utils.Min(right-i, lpc[mirror])
		}
		for sb[i+1+lpc[i]] == sb[i-(1+lpc[i])] {
			lpc[i]++
		}
		if i+lpc[i] > right {
			center = i
			right = i + lpc[i]
		}
	}
	var maxLen int
	for i := 1; i < nsb; i++ {
		if lpc[i] > maxLen {
			maxLen = lpc[i]
			center = i
		}
	}
	return s[(center-1-maxLen)/2 : (center-1+maxLen)/2]

}
