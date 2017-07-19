package excel

func ConvertToTitle(n int) string {
	var ret string
	for n != 0 {
		ret = string('A'+(n-1)%26) + ret
		n = (n - 1) / 26
	}
	return ret
}
