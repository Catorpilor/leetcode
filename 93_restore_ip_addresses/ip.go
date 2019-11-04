package ip

func restoreIpAddresses(s string) []string {
	var res []string
	n := len(s)
	if n < 4 {
		return res
	}
	permute(&res, []string{}, "", s)
	return res
}

func permute(res *[]string, parts []string, bid, s string) {
	if len(s) == 0 {
		// found a valid ip
		tmp := ""
		for idx, s := range parts {
			if idx == 0 {
				tmp = s
			} else {
				tmp = tmp + "." + s
			}
		}
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(s)-3+len(parts); i++ {
		bid += string(s[i])

	}
}
