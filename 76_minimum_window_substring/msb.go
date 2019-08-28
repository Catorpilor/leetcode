package msb

import "fmt"

// MinWindow returns the minimum substring that contains t
func MinWindow(s, t string) string {
	if len(t) < 1 {
		return ""
	}
	cache := make(map[byte]int, len(t))
	for i := range t {
		cache[t[i]]++
	}
	var beg, end, head int
	curLen := 1 << 31
	count := len(cache)
	for end < len(s) {
		fmt.Printf("end: %d, s[end]:%s\n", end, string(s[end]))
		if v, exists := cache[s[end]]; exists {
			v--
			cache[s[end]] = v
			if v == 0 {
				count--
			}
			fmt.Printf("cache[s[end]]=%d, count: %d\n", v, count)
		}
		fmt.Printf("cache now: %v\n", cache)
		end++
		for count == 0 {
			fmt.Printf("beg: %d, s[beg]:%s\n", beg, string(s[beg]))
			if v, exists := cache[s[beg]]; exists {
				v++
				cache[s[beg]] = v
				if v > 0 {
					count++
				}
				fmt.Printf("cache[s[beg]]=%d, count: %d\n", v, count)
			}
			if end-beg < curLen {
				curLen = end - beg
				head = beg
			}
			beg++
		}
	}
	if curLen == 1<<31 {
		return ""
	}
	return s[head : head+curLen]
}
