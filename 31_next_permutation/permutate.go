package permuate

import "sort"

/*
* 1st, find the largest i that nums[i:] is non-incresing(week descreasing)
* 2nd, choose the pivot at i-1 pos, and swap  it with the mininum one in the
* suffix which is greater than nums[i-1]
* so we get the minimal increase of nums[:i]
* 3rd. reverse nums[i:] to compute the next permutation.
* for example
* idxs=[0,1,2,3,4,5,6]
* nums=[0,1,2,5,3,3,0]
* first we got i=3 and choose pivot = 2
* then find  the minimal elements in suffix that is greater than pivot which j = 5
* swap(2,5) becomes
* nums=[0,1,3,5,3,2,0]
* then reverse the suffix to get next permution
* [0,1,3,0,2,3,5]
 */

func NextPermutation(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	i := n - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}
	if i <= 0 {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] <= nums[j]
		})
		return nums
	}
	j := n - 1
	for nums[j] <= nums[i-1] {
		j--
	}
	nums[i-1], nums[j] = nums[j], nums[i-1]
	j = n - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i, j = i+1, j-1
	}
	return nums
}
