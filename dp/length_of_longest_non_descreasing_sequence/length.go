package length

func Lond(input []int) int {
	l := len(input)
	if l <= 1 {
		return l
	}
	s := make([]int, l)
	for i := 0; i < len(s); i++ {
		s[i] = 1
	}

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if input[i] >= input[j] && s[j]+1 > s[i] {
				s[i] = s[j] + 1
			}
		}
	}
	return s[l-1]
}
