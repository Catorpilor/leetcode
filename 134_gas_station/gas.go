package gas

func canComplete(gas, costs []int) int {
	return useStraight(gas, costs)
}

// useStraight time complexity O(N), space complexity O(1)
func useStraight(gas, cost []int) int {
	// if a car starts at a and can not reach b, that any station
	// between a and b can not reach b
	// if the total number of gas is greater than the total number
	// of cost there must be a solution
	var start, tank, total int
	for i := 0; i < len(gas); i++ {
		tank = tank + gas[i] - cost[i]
		if tank < 0 {
			total += tank
			start = i + 1
			tank = 0
		}
	}
	if total+tank < 0 {
		return -1
	}
	return start
}

// useTwoPasses time complexity O(N), space complexity O(1)
func useTwoPasses(gas, cost []int) int {
	n := len(gas)
	dsum := 0
	for i := range gas {
		dsum += gas[i] - cost[i]
	}
	if dsum < 0 {
		return -1
	}
	var start, tank int
	for i := 0; i < n; i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	return start
}
