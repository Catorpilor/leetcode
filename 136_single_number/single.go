package single

func singleNum(arr []int) int {
	var ret int
	for _, a := range arr {
		ret ^= a
	}
	return ret
}
