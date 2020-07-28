package bulb

func minFlips(target string) int {
	return useGreedy(target)
}

// useBruteForce time complexity O(N), space complexity O(1)
func useGreedy(target string) int {
	var ans int
	n := len(target)
	if n < 1 {
		return ans
	}
	var first int
	for {
		for ; first < n; first++ {
			if target[first] == '0' {
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

// simple is a simpler approach for useGreedy, just use one loop.
// time complexity O(N), space complexity O(1)
func simple(target string) int {
	var ans int
	start := '0'
	for _, c := range target {
		if c != start {
			ans++
			if start == '0' {
				start = '1'
			} else {
				start = '0'
			}
		}
	}
	return ans
}
