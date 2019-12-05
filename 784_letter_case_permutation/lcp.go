package lcp

func letterCasePermutations(s string) []string {
	var res []string
	helper(&res, []byte(s), 0)
	return res
}

func helper(res *[]string, bs []byte, idx int) {
	if len(bs) == idx {
		tmp := make([]byte, len(bs))
		copy(tmp, bs)
		*res = append(*res, string(tmp))
		return
	}
	if bs[idx] >= '0' && bs[idx] <= '9' {
		helper(res, bs, idx+1)
		return
	}
	bs[idx] = toUpper(bs[idx])
	helper(res, bs, idx+1)
	bs[idx] = toLower(bs[idx])
	helper(res, bs, idx+1)
}

func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		c = c - 'a' + 'A'
	}
	return c
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		c = c + 'a' - 'A'
	}
	return c
}
