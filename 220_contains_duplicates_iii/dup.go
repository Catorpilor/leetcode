package dup

import "math"

func containsDups(arr []int, k, t int) bool {
	return useBucket(arr, k, t)
}

// useBucket time complexity O(N), space complexity O(N)
func useBucket(arr []int, k, t int) bool {
	if k < 1 || t < 0 {
		return false
	}
	n := len(arr)
	set := make(map[int64]int64, n)
	for i := range arr {
		repNum := int64(arr[i]) - math.MinInt32
		bkt := repNum / (int64(t) + 1)
		if _, exists := set[bkt]; exists {
			return true
		}
		if v, exists := set[bkt-1]; exists {
			if repNum-v <= int64(t) {
				return true
			}
		}
		if v, exists := set[bkt+1]; exists {
			if v-repNum <= int64(t) {
				return true
			}
		}
		if i >= k {
			lastBkt := (int64(arr[i-k]) - math.MinInt32) / (int64(t) + 1)
			delete(set, lastBkt)
		}
		set[bkt] = repNum
	}
	return false
}
