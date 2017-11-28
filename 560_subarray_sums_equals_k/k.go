package sum

func SubArraySum(nums []int, k int) int {
	// O(n^2)
	n := len(nums)
	if n < 1 {
		return 0
	}
	var ret int
	for i := range nums {
		cur := nums[i]
		if cur == k {
			ret++
		}

		for j := i + 1; j < n; j++ {
			cur += nums[j]
			if cur == k {
				ret++
			}
		}
	}
	return ret
}

func SubArraySum2(nums []int, k int) int {
	var count, sum int
	// record[k]=v means how many sub-arrays sums equals to k
	// preSum for nums [-3,1,2,-3,5] is
	// preSum = [0,-3,-2,0,-3,2]
	// so to get SUM[1,3] which is [1,2,-3] we can use preSum to caculcate which is
	// preSum[4] - preSum[1] = -3 + 3 = 0
	// so if a subarray i to j sums equal to k , which means
	// preSum[j+1] - preSum[i] = k
	record := make(map[int]int)
	record[0] = 1
	for _, n := range nums {
		sum += n
		if v, ok := record[sum-k]; ok {
			count += v
		}
		record[sum]++
	}
	return count
}
