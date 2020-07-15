package angle

import "math"

func angleClock(hour, minutes int) float64 {
	return useMath(hour, minutes)
}

// useMath time complexity O(1) ,space complexity O(1)
func useMath(hour, minutes int) float64 {
	// 1 minute => 6
	// 5/60 = x/minutes
	// all start from clock-wise 0
	// sameSide := false
	// if (hour <= 6 && minutes <= 30) || (hour > 6 && minutes > 30) {
	// sameSide = true
	// }
	hourDiff := float64(minutes) / float64(2)
	hourAngle := float64(hour*30) + hourDiff
	minuteAngle := float64(minutes * 6)
	diff := math.Abs(hourAngle - minuteAngle)
	if diff <= 180.0 {
		return diff
	}
	return float64(360) - diff
}
