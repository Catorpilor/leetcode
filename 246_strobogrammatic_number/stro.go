package stro

func IsStro(num string) bool {
	hm := map[byte]byte{
		'0': '0',
		'1': '1',
		'8': '8',
		'6': '9',
		'9': '6',
	}
	for i, j := 0, len(num)-1; i <= j; i, j = i+1, j-1 {
		if v, ok := hm[num[i]]; !ok {
			return false
		} else {
			if v != num[j] {
				return false
			}
		}
	}
	return true
}
