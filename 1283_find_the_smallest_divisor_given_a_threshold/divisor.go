package divisor

func smallestDivisor(nums []int, threshold int) int {
	return useBinarySearch(nums, threshold)
}

// useBinarySearch time complexity O(N*logM) where M is pSum+1, space complexity O(1)
func useBinarySearch(nums []int, threshold int) int {
	psum := 0
	for i := range nums {
		psum += nums[i]
	}
	// l, r := n              ------> prefixSum[n-1]
	// divisor prefixSum[n-1] ------> 1
	ld, rd := 1, psum+1 // psum+1  here for [1] and threshold is 1
	var ans int
	for ld < rd {
		mid := ld + (rd-ld)>>1
		if cal(nums, mid) <= threshold {
			// we find a winner here but can we do better
			ans = mid
			rd = mid
		} else {
			ld = mid + 1
		}
	}
	return ans
}

func cal(nums []int, divisor int) int {
	var ans int
	for i := range nums {
		/*
			if nums[i] <= divisor {
				ans++
			} else {
				ans += nums[i] / divisor
				if nums[i]&(divisor-1) != 0 {
					ans++
				}
			}
		*/
		ans += (nums[i] + divisor - 1) / divisor
	}
	return ans
}
