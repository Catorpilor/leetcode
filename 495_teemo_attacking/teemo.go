package teemo

func teemoAttack(timeSeries []int, duration int) int {
	return useStraight(timeSeries, duration)
}

// useStraight time complexity O(N), space complexity O(1)
func useStraight(ts []int, duration int) int {
	n := len(ts)
	if n < 1 {
		return 0
	}
	var ans int
	pre := ts[0]
	for i := 1; i < n; i++ {
		diff := ts[i] - pre
		if diff > duration {
			ans += duration
			pre = ts[i]
		} else {
			ans += diff
		}
	}
	ans += duration
	return ans
}
