package anagrams

func FindAnagrams(s, p string) []int {
	var ret []int
	if len(s) < len(p) {
		return ret
	}
	var target, window [26]int
	for i, c := range p {
		target[c-'a'] += 1
		window[s[i]-'a'] += 1
	}
	if target == window {
		ret = append(ret, 0)
	}
	for i := 1; i <= len(s)-len(p); i++ {
		window[s[i-1]-'a'] -= 1
		window[s[i+len(p)-1]-'a'] += 1
		if target == window {
			ret = append(ret, i)
		}
	}
	return ret
}

func findAnagrams2(s string, p string) []int {
	var ret []int
	if len(s) < len(p) {
		return ret
	}

	for i := 0; i <= len(s)-len(p); i++ {
		if isAnagram(s[i:i+len(p)], p) {
			ret = append(ret, i)
		}
	}
	return ret
}

func isAnagram(s, t string) bool {
	h := [26]int{}
	for i := 0; i < len(s); i++ {
		h[s[i]-'a'] += 1
		h[t[i]-'a'] -= 1
	}
	for _, v := range h {
		if v != 0 {
			return false
		}
	}
	return true
}
