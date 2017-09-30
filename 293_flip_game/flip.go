package flip

func Flip(s string) []string {
	ret := make([]string, 0, len(s))
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '+' && s[i+1] == '+' {
			str := s[:i] + "--"
			if i+2 < len(s) {
				str += s[i+2:]
			}
			ret = append(ret, str)
		}
	}
	return ret
}
