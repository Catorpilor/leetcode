package ip

import (
	"strconv"
	"strings"
)

func validateIP(ip string) string {
	return useDividAndConqure(ip)
}

// useDividAndConqure time complexity O(N), space complexity O(1)
func useDividAndConqure(ip string) string {
	if strings.Count(ip, ".") == 3 {
		return isValidIPV4(ip)
	}
	if strings.Count(ip, ":") == 7 {
		return isValidIPV6(ip)
	}
	return "Neither"
}

func isValidIPV4(ip string) string {
	v4Splitter := func(c rune) bool {
		return c == '.'
	}
	segs := strings.FieldsFunc(ip, v4Splitter)
	if len(segs) != 4 {
		return "Neither"
	}
	for _, seg := range segs {
		num, err := strconv.Atoi(seg)
		if err != nil || num < 0 || num > 255 || strconv.FormatInt(int64(num), 10) != seg {
			return "Neither"
		}
	}
	return "IPV4"
}

func isValidIPV6(ip string) string {
	v6Splitter := func(c rune) bool {
		return c == ':'
	}
	segs := strings.FieldsFunc(ip, v6Splitter)
	if len(segs) != 8 {
		return "Neither"
	}
	for _, seg := range segs {
		if len(seg) > 4 || len(seg) < 1 {
			return "Neither"
		}
		num, err := strconv.ParseInt(seg, 16, 64)
		// fmt.Printf("parse seg:%s num:%d, err:%v\n", seg, num, err)
		if err != nil || num < 0 || seg[0] == '-' {
			return "Neither"
		}
	}
	return "IPV6"
}
