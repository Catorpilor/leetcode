package record

func CheckRecords(s string) bool {
	cnt := [3]int{}
	// A -> 0
	// L -> 1
	// P -> 2
	for i, c := range s {
		switch c {
		case 'A':
			cnt[0] += 1
			if cnt[0] > 1 {
				return false
			}
		case 'L':
			cnt[1] += 1
			if cnt[1] > 2 && s[i-1] == 'L' && s[i-2] == 'L' {
				return false
			}
		default:
			cnt[2] += 1
		}
	}
	return true
}
