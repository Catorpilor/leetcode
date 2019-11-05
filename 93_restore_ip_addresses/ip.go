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

func iterator(s string) []string {
	var res []string
	n := len(s)
	if n < 4 {
		return res
	}
	for a := 1; a < 4; a++ {
		for b := 1; b < 4; b++ {
			for c := 1; c < 4; c++ {
				for d := 1; d < 4; d++ {
					if a+b+c+d == n {
						partA := s[:a]
						partB := s[a : a+b]
						partC := s[b+a : c+a+b]
						partD := s[c+a+b:]
						ai, _ := strconv.Atoi(partA)
						bi, _ := strconv.Atoi(partB)
						ci, _ := strconv.Atoi(partC)
						di, _ := strconv.Atoi(partD)
						if ai <= 255 && bi <= 255 && ci <= 255 && di <= 255 {
							if len(partA) > 0 && partA[0] == '0' {
								continue
							}
							res = append(res, partA+"."+partB+"."+partC+"."+partD)
						}
					}
				}
			}
		}
	}
	return res
}
