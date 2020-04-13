package stone

import "sort"

func lastWeightStones(stones []int) int {
	return useSort(stones)
}

// useSort time complexity O(N^2lgN), space complexity O(1)
func useSort(stones []int) int {
	n := len(stones)
	sort.Ints(stones)
	i := n - 2
	for i >= 0 && i+1 < len(stones) {
		if stones[i] == stones[i+1] {
			stones = stones[:i]
			i -= 2
		} else {
			stones[i] = stones[i+1] - stones[i]
			stones = stones[:i+1]
			sort.Ints(stones)
			i--
		}
	}
	if len(stones) > 0 {
		return stones[0]
	}
	return 0
}

// useBucket time complexity O(N+W), space compleixty O(W)
func useBucket(stones []int) int {
	n := len(stones)
	if n < 1 {
		return 0
	}
	maxWeight := stones[0]
	for i := 1; i < n; i++ {
		if stones[i] > maxWeight {
			maxWeight = stones[i]
		}
	}
	bucket := make([]int, maxWeight+1)
	// distrubute to buckets
	for i := range stones {
		bucket[stones[i]]++
	}
	currentWeight := maxWeight
	biggestWeight := 0
	for currentWeight > 0 {
		if bucket[currentWeight] == 0 {
			currentWeight--
		} else if biggestWeight == 0 {
			bucket[currentWeight] %= 2
			if bucket[currentWeight] == 1 {
				biggestWeight = currentWeight
			}
			currentWeight--
		} else {
			bucket[currentWeight]--
			if biggestWeight-currentWeight <= currentWeight {
				bucket[biggestWeight-currentWeight]++
				biggestWeight = 0
			} else {
				biggestWeight -= currentWeight
			}
		}
	}
	return biggestWeight
}
