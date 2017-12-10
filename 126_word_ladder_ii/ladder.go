package ladder

func Ladders(start, end string, ws []string) [][]string {
	var ret [][]string
	temp := []string{start}
	dict := make(map[string]bool) // maps ws for quick check
	for i := range ws {
		dict[ws[i]] = true
	}
	helper(&ret, ws, &temp, start, end)
	return ret
}

func helper(res *[][]string, ws []string, temp *[]string, pivotS, es string) {
	if pivotS == es {
		if *res != nil && len(*temp) > len((*res)[0]) {
			return
		}
		// if *res == nil || *res != nil && len((*res)[0]) == len(*temp) {
		// 	bak := make([]string, len(*temp))
		// 	copy(bak, *temp)
		// 	*res = append(*res, bak)
		// }
		if *res != nil && len(*temp) < len((*res)[0]) {
			*res = nil
			// bak := make([]string, len(*temp))
			// copy(bak, *temp)
			// *res = append(*res, bak)
			// return
		}
		// if len((*res)[0]) == len(*temp) {
		// 	bak := make([]string, len(*temp))
		// 	copy(bak, *temp)
		// 	*res = append(*res, bak)
		// }
		bak := make([]string, len(*temp))
		copy(bak, *temp)
		*res = append(*res, bak)
		// fmt.Println(bak)

	} else {
		// find next pivotS
		for i := range ws {
			if !exists(*temp, ws[i]) && dis(pivotS, ws[i]) == 1 {
				// fmt.Println(*temp)
				n := len(*temp)
				*temp = append(*temp, ws[i])
				helper(res, ws, temp, ws[i], es)
				*temp = (*temp)[:n]
			}
		}
	}
}

func exists(ws []string, p string) bool {
	for i := range ws {
		if ws[i] == p {
			return true
		}
	}
	return false
}

func dis(s, t string) int {
	// assume s and t has the same length
	var ret int
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			ret++
		}
	}
	return ret
}

// func distance(s string, dict map[string]bool) []string{

// }
