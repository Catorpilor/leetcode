package wb

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
