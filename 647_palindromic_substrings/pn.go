package pn

func countSubString(s string) int {
	return expendFromCenter(s)
}

// expendFromCenter time compelxity O(N^2), space complexity O(1)
func expendFromCenter(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var count int
	for i := 0; i < len(s); i++ {
		checkPanlidrome(s, i, i, &count)   // odd ones, middle is i
		checkPanlidrome(s, i, i+1, &count) // even ones middle is in [i,i+1]
	}
	return count
}

func checkPanlidrome(s string, l, r int, count *int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		*count, l, r = *count+1, l-1, r+1
	}
}

// useBruteForce time complexity O(N^3), space complexity O(1)
func useBruteForce(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	ans := n
	for k := 2; k <= n; k++ {
		for i := 0; i+k <= n; i++ {
			if isPalindrome(s[i : i+k]) { // O(N)
				ans++
			}
		}
	}
	return ans
}

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

// expendFromCenterWithSkip time complexity O(N^2), space complexity O(1)
func expendFromCenterWithSkip(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	var idx, ans int
	for idx < n {
		front := idx
		for idx < n && s[idx] == s[front] {
			// skip all the identical characters
			idx++
		}
		skipped := idx - front
		// use formular (n+1)*n/2
		ans += (skipped + 1) * skipped / 2
		// mid in [front-1, idx]
		ans += extend(s, front-1, idx)
	}
	return ans
}

func extend(s string, l, r int) int {
	var res int
	n := len(s)
	for l >= 0 && r < n && s[l] == s[r] {
		res++
		l--
		r++
	}
	return res
}
