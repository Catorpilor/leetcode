package mgm

func minMutation(beg, end string, bank []string) int {
	return bfsWithPattern(beg, end, bank)
}

func bfsWithPattern(beg, end string, bank []string) int {
	hset := make(map[string][]string, len(bank))
	visited := make(map[string]bool, len(bank))
	leveld := make([]string, 0, len(bank))
	pattern(hset, bank, len(beg))
	for i := 0; i < len(beg); i++ {
		key := beg[:i] + "*" + beg[i+1:]
		leveld = append(leveld, hset[key]...)
	}
	visited[beg] = true
	level := 1
	for len(leveld) != 0 {
		next := make([]string, 0, len(bank))
		for _, str := range leveld {
			if str == end {
				return level
			}
			if visited[str] {
				continue
			}
			visited[str] = true
			for i := 0; i < len(beg); i++ {
				key := str[:i] + "*" + str[i+1:]
				next = append(next, hset[key]...)
			}
		}
		leveld = next
		level++
	}
	return -1
}

func pattern(hset map[string][]string, bank []string, l int) {
	for _, str := range bank {
		for i := 0; i < l; i++ {
			key := str[:i] + "*" + str[i+1:]
			hset[key] = append(hset[key], str)
		}
	}
}
