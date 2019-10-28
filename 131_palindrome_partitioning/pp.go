package pp

import "fmt"

func partition(s string) [][]string {
	var res [][]string
	n := len(s)
	for i := 0; i < n; i++ {
		permute(&res, []string{}, i, i, n, s)
	}
	return res
}

func permute(res *[][]string, bid []string, start, cur, n int, s string) {
	fmt.Printf("bid: %v, start: %d, cur: %d\n", bid, start, cur)
	if cur == n {
		tmp := make([]string, len(bid))
		copy(tmp, bid)
		*res = append(*res, tmp)
		return
	}
	// a 0, 1
	// ab
	for i := cur; i < n; i++ {
		if isPalindrome(s[start : i+1]) {
			fmt.Printf("bid: %v, s[start:i+1] : %s, i: %d\n", bid, s[start:i+1], i)
			bid = append(bid, s[start:i+1])
		}
		permute(res, bid, start, i+1, n, s)
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
