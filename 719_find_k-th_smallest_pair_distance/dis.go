package dis

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func KthDistance(nums []int, k int) int {
	n := len(nums)
	if n <= 1 || k < 1 || k > n*(n-1)/2 {
		return -1
	}
	// brute force
	// got TLE
	distances := make([]int, 0, n*(n-1)/2+1)
	distances = append(distances, 0)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			distances = append(distances, utils.Abs(nums[i]-nums[j]))
		}
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i] <= distances[j]
	})
	return distances[k]
}

func KthDistance2(nums []int, k int) int {
	n := len(nums)
	if n <= 1 || k < 1 || k > n*(n-1)/2 {
		return -1
	}
	// since nums[i] in range [0,1000000)
	// so the maxdiff for pair(i,j) would be less than 1000000
	var hm [1000000]int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			hm[utils.Abs(nums[i]-nums[j])]++
		}
	}
	for i := range hm {
		if hm[i] >= k {
			return i
		}
		k -= hm[i]
	}
	return -1
}

func KthDistance3(nums []int, k int) int {
	n := len(nums)
	if n <= 1 || k < 1 || k > n*(n-1)/2 {
		return -1
	}
	// sort nums
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	low := nums[1] - nums[0]
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] < low {
			low = nums[i] - nums[i-1]
		}
	}
	hi := nums[n-1] - nums[0]
	var mid int
	for low < hi {
		mid = low + (hi-low)/2
		// count how many pairs in nums with distance <= mid
		if helper(nums, mid) < k {
			low = mid + 1
		} else {
			hi = mid
		}
	}
	return low
}

func helper(nums []int, v int) int {
	var ret, j int
	n := len(nums)
	for i := 0; i < n-1; i++ {
		j = i + 1
		for j < n && nums[j]-nums[i] <= v {
			j++
		}
		ret += j - i - 1
	}
	return ret
}

func upper(nums []int, low, hi, v int) int {
	if nums[hi] <= v {
		return hi + 1
	}
	var mid int
	for low < hi {
		mid = low + (hi-low)/2
		if nums[mid] <= v {
			low = mid + 1
		} else {
			hi = mid
		}
	}
	return low
}

func helper2(nums []int, v int) int {
	var ret int
	n := len(nums)
	for i := 0; i < n-1; i++ {
		ret += upper(nums, i, n-1, nums[i]+v) - i - 1
	}
	return ret
}

func KthDistance4(nums []int, k int) int {
	n := len(nums)
	if n <= 1 || k < 1 || k > n*(n-1)/2 {
		return -1
	}
	// sort nums
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	low := nums[1] - nums[0]
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] < low {
			low = nums[i] - nums[i-1]
		}
	}
	hi := nums[n-1] - nums[0]
	var mid int
	for low < hi {
		mid = low + (hi-low)/2
		// count how many pairs in nums with distance <= mid
		if helper2(nums, mid) < k {
			low = mid + 1
		} else {
			hi = mid
		}
	}
	return low
}
