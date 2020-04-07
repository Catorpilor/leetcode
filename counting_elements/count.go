package count

func countElements(arr []int) int {
	return useHashmap(arr)
}

func useHashmap(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}
	set := make(map[int]bool, n)
	for _, num := range arr {
		set[num] = true
	}
	var ans int
	for _, num := range arr {
		if set[num+1] {
			ans++
		}
	}
	return ans
}
