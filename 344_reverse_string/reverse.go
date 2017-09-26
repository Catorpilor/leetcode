package reverse

func Reverse(s string) string {
	if len(s) < 2 {
		return s
	}
	ret := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		ret = append(ret, byte(s[i]))
	}
	return string(ret)
}
