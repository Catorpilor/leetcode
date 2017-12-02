package dis

func Rearrange(s string, k int) string {
	hm := make([]int, 26)
	valid := make([]int, 26)
	var id, maxF, numOfMaxF int
	for _, c := range s {
		id = int(c - 'a')
		hm[id]++
		if hm[id] > maxF {
			maxF, numOfMaxF = hm[id], 1
		} else if hm[id] == maxF {
			numOfMaxF++
		}
	}
	if (maxF-1)*k+numOfMaxF > len(s) {
		return ""
	}
	ret := make([]byte, len(s))
	var pos int
	for i := range s {
		pos = findValidMax(hm, valid, i)
		// if pos == -1 {
		// 	return ""
		// }
		hm[pos]--
		valid[pos] = i + k
		ret[i] = byte(pos) + 'a'
	}
	return string(ret)
}

func findValidMax(count []int, valid []int, idx int) int {
	curMax, pos := 0, -1
	for i := range count {
		if count[i] > curMax && idx >= valid[i] {
			curMax = count[i]
			pos = i
		}
	}
	return pos
}
