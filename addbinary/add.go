package addbinary

func reverse(runes []rune) string {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func AddBinary(a, b string) string {
	carry := 0
	i := len(a) - 1
	j := len(b) - 1
	l := i
	if j > l {
		l = j
	}
	ret := make([]rune, 0, l)
	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(b[j] - '0')
			j--
		}
		ret = append(ret, rune((carry&1)+'0'))
		carry /= 2
	}
	return reverse(ret)
}
