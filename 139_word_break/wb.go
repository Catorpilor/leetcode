package wb

import "github.com/catorpilor/leetcode/utils"

func wordBreak(s string, words []string) bool {
	n := len(s)
	nw := len(words)
	if n == 0 || nw == 0 {
		return false
	}
	// make a cache
	cache := make(map[string]bool)
	for _, word := range words {
		cache[word] = true
	}
	// set stores the index that can not find a match in dict
	set := make(map[int]bool)
	return dfs(cache, s, 0, n, set)
}

// time complexity O(2^n)
func dfs_tle(cache map[byte][]string, s string, idx, n int) bool {
	if idx == n {
		return true
	}
	words := cache[s[idx]]
	if len(words) == 0 {
		return false
	}
	var ret bool
	var upper int
	for _, word := range words {
		upper = idx + len(word)
		if upper <= n && s[idx:upper] == word {
			ret = dfs_tle(cache, s, upper, n)
		}
	}
	return ret
}

func dfs(cache map[string]bool, s string, idx, n int, set map[int]bool) bool {
	if idx == n {
		return true
	}
	if set[idx] {
		return false
	}
	for i := idx + 1; i <= n; i++ {
		if cache[s[idx:i]] {
			if dfs(cache, s, i, n, set) {
				return true
			}
			set[i] = true
		}
	}
	set[idx] = true
	return false
}

func wordBreak2(s string, words []string) bool {
	n := len(s)
	nw := len(words)
	if n == 0 || nw == 0 {
		return false
	}
	set := make(map[string]bool, len(words))
	for _, word := range words {
		set[word] = true
	}

	return bfs(s, set, n)
}

func bfs(s string, set map[string]bool, n int) bool {
	queue := utils.NewQueue()
	queue.Enroll(0)
	visited := make([]bool, n)
	for !queue.IsEmpty() {
		start := queue.Pull().(int)
		if !visited[start] {
			for end := start + 1; end <= n; end++ {
				if set[s[start:end]] {
					queue.Enroll(end)
					if end == n {
						return true
					}
				}
			}
		}
		visited[start] = true
	}
	return false
}

func wordBreak3(s string, words []string) bool {
	n := len(s)
	nw := len(words)
	set := make(map[string]bool, nw)
	for _, word := range words {
		set[word] = true
	}

	return useDp(s, set, n)
}

func useDp(s string, set map[string]bool, n int) bool {
	dp := make([]bool, n+1)
	dp[0] = true
	// i represents the length of substring s' starting from beginning.
	// j refers to the index partition the current substring s' into smaller substrings s1'[0:j] and s2'[j:i]
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && set[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
