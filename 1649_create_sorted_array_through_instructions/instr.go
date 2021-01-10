package inster

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func createSortedArray(ins []int) int {
	return useBF(ins)
}

// useBF time complexity O(N^2*log(N)), space complexity O(N)
func useBF(instructions []int) int {
	n := len(instructions)
	nums := make([]int, 0, n)
	var ans int
	for _, num := range instructions {
		ans += cost(nums, num)
		nums = append(nums, num)
		sort.Ints(nums)
	}
	return ans
}

func cost(nums []int, target int) int {
	if len(nums) == 0 || nums[0] >= target || nums[len(nums)-1] <= target {
		return 0
	}
	return utils.Min(lb(nums, target), up(nums, target))
}

// lb lower_bound find the proper position for target, time complexity O(log(N))
func lb(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// fmt.Printf("lower_bound nums:%v with target: %d return: %d\n", nums, target, l)
	return l
}

// up upper_bound find the proper position for target, time complexity O(log(N))
func up(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := r - (r-l)>>1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	// fmt.Printf("upper_bound nums:%v with target: %d return: %d\n", nums, target, l)
	return len(nums) - l - 1
}
