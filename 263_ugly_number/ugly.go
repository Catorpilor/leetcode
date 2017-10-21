package ugly

func UglyNumber(n int) bool {
	for i := 2; i < 6; i++ {
		for n%i == 0 {
			n /= i
		}
	}
	return n == 1
}
