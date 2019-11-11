package wb

func wordBreak(s string, words []string) bool {
	n := len(s)
	nw := len(words)
	if n == 0 || nw == 0 {
		return false
	}
	// make a cache
	cache := make(map[byte][]string)
	for _, word := range words {
		cache[word[0]] = append(cache[word[0]], word)
	}
	return dfs(cache, s, 0, n)
}

func dfs(cache map[byte][]string, s string, idx, n int) bool {
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
			ret = dfs(cache, s, upper, n)
		}
	}
	return ret
}
