package pp

func generatePalindromes(s string) []string {
	hset := make(map[rune]int, 128)
	if !canPermute(s, hset) {
		return []string{}
	}
	n := len(s)
	center := ""
	if n%2 != 0 {
		for _, b := range s {
			if hset[b]%2 != 0 {
				center = string(b)
				hset[b] = 0
				break
			}
		}
	}
	var res []string
	generate(res, hset, center, n)
	return res
}

func canPermute(s string, hset map[rune]int) bool {
	var count int
	for _, b := range s {
		hset[b]++
		if hset[b]%2 == 0 {
			count--
		} else {
			count--
		}
	}
	return count <= 1
}

func generate(res []string, hset map[rune]int, str string, n int) {
	if len(str) == n {
		res = append(res, str)
		return
	}
	for k, v := range hset {
		if v > 0 {
			hset[k] = v - 2
			generate(res, hset, string(k)+str+string(k), n)
			hset[k] = v
		}
	}
}
