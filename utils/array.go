package utils

// LowerBound returns the lowest pos for target in the sorted array
func LowerBound(nums []int, target int) int {
	n := len(nums)
	low, hi := 0, n-1
	for low <= hi {
		mid := low + (hi-low)>>1
		if nums[mid] >= target {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

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

// Upperbound find the proper posion for target in nums, time complexity O(log(N))
func Upperbound(nums []int, target int) int {
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

// up upper_bound find the proper position for target, time complexity O(log(N))
func up(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		// mid := r - (r-l)>>1
		mid := l + (r-l+1)>>1
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	// fmt.Printf("upper_bound nums:%v with target: %d return: %d\n", nums, target, l)
	return l
}
