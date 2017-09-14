package dif

func Dif(s, t string) byte {
	hm := make(map[rune]int)
	for _, c := range s {
		if _, ok := hm[c]; ok {
			hm[c] += 1
		} else {
			hm[c] = 1
		}
	}
	var ret byte
	for _, k := range t {
		if _, ok := hm[k]; ok {
			if hm[k] == 1 {
				delete(hm, k)
			} else {
				hm[k] -= 1
			}
		} else {
			ret = byte(k)
			break
		}

	}
	return ret
}

func Dif2(s, t string) byte {
	var ret int
	for _, c := range s {
		ret ^= int(c)
	}
	for _, v := range t {
		ret ^= int(v)
	}
	return byte(ret)
}
