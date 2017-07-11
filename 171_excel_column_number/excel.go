package excel

func TitleToNumber(s string) int {
	var ret int
	for i := 0; i < len(s); i++ {
		ret *= 26
		ret += int(s[i]-'A') + 1
	}
	return ret
}
