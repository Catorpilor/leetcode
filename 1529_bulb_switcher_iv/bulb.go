package bulb

func minFlips(target string) int {
	return useBruteForce(target)
}


// useBruteForce time complexity O(N), space complexity O(1)
func useBruteForce(target string) int {
	var ans int
	n := len(target)
	if n < 1 {
		return ans
	}
	var first int
	for {
		for ; first < n; first++ {
			if target[first] == '0'{
				continue
			}
			ans++
			break
		}
		first++
		if first >= n {
			break
		}
		for ; first < n; first++ {
			if target[first] == '1' {
				continue
			}
			ans++
			break
		}
	}
	return ans
}
