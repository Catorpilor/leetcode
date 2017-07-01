package strstr

func StrStr(haystack, needle string) int {
	if haystack == needle {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	for i := range haystack {
		if (i+len(needle) <= len(haystack)) && (haystack[i:i+len(needle)] == needle) {
			return i
		}
	}
	return -1
}
