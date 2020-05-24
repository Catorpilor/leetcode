package vowels

import (
	"strings"

	"github.com/catorpilor/leetcode/utils"
)

func maxVowels(s string, k int) int {
	return useSlideWindow(s, k)
}

// useSlideWindow time complexity O(N), space complexity O()
func useSlideWindow(s string, k int) int {
	n := len(s)
	if n < 1 {
		return 0
	}
	if k == 1 {
		return 1
	}
	var cur int
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			cur++
		}
	}
	ans := cur
	for j := k; j < n; j++ {
		if isVowel(s[j]) {
			cur++
		}
		if isVowel(s[j-k]) {
			cur--
		}
		ans = utils.Max(ans, cur)
	}
	return ans
}

func isVowel(b byte) bool {
	return strings.IndexByte("aeiou", b) != -1
}
