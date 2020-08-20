package balls

import "sort"

func maxDistance(position []int, m int) int {
	return useBinarySearch(position, m)
}

// useBinarySearch time complexity O(N*logM) where M = max(position) - min(position), space complexity O(1)
func useBinarySearch(position []int, m int) int {
	sort.Ints(position)
	n := len(position)
	if m == 2 {
		return position[n-1] - position[0]
	}
	// l, r means the min and max distance between two balls
	// ans its monotonically increasing, so we could use binary search
	l, r := 0, position[n-1]-position[0]
	for l < r {
		mid := r - (r-l)/2
		if count(mid, position) >= m {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l

}

func count(d int, position []int) int {
	ans, cur := 1, position[0]
	for i := 1; i < len(position); i++ {
		if position[i]-cur >= d {
			ans++
			cur = position[i]
		}
	}
	return ans
}
