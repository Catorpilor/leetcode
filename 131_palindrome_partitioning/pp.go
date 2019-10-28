package pp

import "fmt"

func partition(s string) [][]string {
	var res [][]string
	permute(&res, []string{}, 0, s)
	return res
}

func permute(res *[][]string, bid []string, cur int, s string) {
	fmt.Printf("bid: %v, cur: %d\n", bid, cur)
	if cur == len(s) {
		tmp := make([]string, len(bid))
		copy(tmp, bid)
		*res = append(*res, tmp)
	} else {
		for i := cur; i < len(s); i++ {
			if isPalindrome(s[cur : i+1]) {
				bl := len(bid)
				bid = append(bid, s[cur:i+1])
				permute(res, bid, i+1, s)
				bid = bid[:bl]
			}
		}
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
