package ip

import "strconv"

func restoreIpAddresses(s string) []string {
	var res []string
	n := len(s)
	if n < 4 {
		return res
	}
	permute(&res, "", s, 0, 0)
	return res
}

func permute(res *[]string, bid, s string, pos, dots int) {
	if len(s)-pos > 3*(4-dots) {
		return
	}
	if dots == 4 && pos == len(s) {
		*res = append(*res, bid)
		return
	}
	for i := 1; i < 4; i++ {
		if pos+i > len(s) {
			return
		}
		seg := s[pos : pos+i]
		si, _ := strconv.Atoi(seg)
		if len(seg) > 1 && seg[0] == '0' || i == 3 && si > 255 {
			continue
		}
		tmp := seg
		if dots != 0 {
			tmp = bid + "." + seg
		}
		permute(res, tmp, s, pos+i, dots+1)
	}
}
