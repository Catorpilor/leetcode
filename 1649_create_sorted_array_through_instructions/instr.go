package inster

import (
	"github.com/catorpilor/leetcode/utils"
)

func createSortedArray(ins []int) int {
	return useBF(ins)
}

const (
	mod = int(1e9 + 7)
)

// useBF time complexity O(N^2*log(N)), space complexity O(N)
func useBF(instructions []int) int {
	n := len(instructions)
	nums := make([]int, 0, n)
	var ans int
	for i, num := range instructions {
		lc, rc := cost(nums, num)
		ans += utils.Min(lc, i-rc)
		nums = append(nums, 0)
		copy(nums[rc+1:], nums[rc:])
		nums[rc] = num
	}
	return ans % mod
}

func cost(nums []int, target int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	return lb(nums, target), upper_bound(nums, target)
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
// func up(nums []int, target int) int {
// 	l, r := 0, len(nums)-1
// 	if l == r && target > nums[l] {
// 		return l + 1
// 	}
// 	for l < r {
// 		mid := r - (r-l)>>1
// 		if nums[mid] <= target {
// 			l = mid
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	fmt.Printf("upper_bound nums:%v with target: %d return: %d\n", nums, target, l)
// 	return l
// }

/*
template<class ForwardIt, class T, class Compare>
ForwardIt upper_bound(ForwardIt first, ForwardIt last, const T& value, Compare comp)
{
    ForwardIt it;
    typename std::iterator_traits<ForwardIt>::difference_type count, step;
    count = std::distance(first, last);

    while (count > 0) {
        it = first;
        step = count / 2;
        std::advance(it, step);
        if (!comp(value, *it)) {
            first = ++it;
            count -= step + 1;
        }
        else
            count = step;
    }
    return first;
}
*/

func upper_bound(nums []int, target int) int {
	count := len(nums)
	l := 0
	for count > 0 {
		step := count >> 1
		mid := l + step
		if nums[mid] <= target {
			l = mid + 1
			count -= step + 1
		} else {
			count = step
		}
	}
	// fmt.Printf("upper_bound nums:%v with target: %d return: %d\n", nums, target, l)
	return l
}
