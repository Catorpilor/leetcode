package median

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func Median(nums1, nums2 []int) float64 {
	// brute force
	// O(m+n) time O(m+n) space
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 && n2 == 0 {
		return 0.0
	}
	merged := make([]int, 0, n1+n2)
	var i, j int
	for i < n1 && j < n2 {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			merged = append(merged, nums2[j])
			j++
		} else {
			merged = append(merged, nums1[i], nums2[j])
			i, j = i+1, j+1
		}
	}
	if i != n1 {
		merged = append(merged, nums1[i:]...)
	}
	if j != n2 {
		merged = append(merged, nums2[j:]...)
	}
	var ret float64
	mid := (n1 + n2) / 2
	if (n1+n2)%2 == 0 {
		// even numbers
		ret = float64(merged[mid-1]+merged[mid]) / float64(2)
	} else {
		ret = float64(merged[mid])
	}
	return ret
}

func Median2(nums1, nums2 []int) float64 {
	// O(lg(min(n1,n2))) time complexity and O(1) space
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 && n2 == 0 {
		return 0.0
	}
	//  binay search on the smaller array
	if n1 > n2 {
		return Median2(nums2, nums1)
	}
	// p1/p2 stands for the partition point for nums1/nums2
	// which has the equation for p1 + p2 = (n1+n2+1) / 2
	// if n1 + n2 is even then we can evenly split the nums1 and nums2
	// if n1 + n2 is odd, then the left split has one more element
	start, end, p1, p2, el := 0, n1, 0, 0, (n1+n2+1)/2
	maxLeft1, maxLeft2, minRight1, minRight2 := math.MinInt32, math.MinInt32, math.MaxInt32, math.MaxInt32
	for start <= end {
		// start = end split the whole array..
		p1 = start + (end-start)/2
		p2 = el - p1
		if p1 != 0 {
			maxLeft1 = nums1[p1-1]
		} else {
			maxLeft1 = math.MinInt32
		}
		if p1 != n1 {
			minRight1 = nums1[p1]
		} else {
			minRight1 = math.MaxInt32
		}
		if p2 != 0 {
			maxLeft2 = nums2[p2-1]
		} else {
			maxLeft2 = math.MinInt32
		}
		if p2 != n2 {
			minRight2 = nums2[p2]
		} else {
			minRight2 = math.MaxInt32
		}
		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (n1+n2)%2 == 0 {
				return float64(utils.Max(maxLeft1, maxLeft2)+utils.Min(minRight1, minRight2)) / float64(2)
			} else {
				return float64(utils.Max(maxLeft1, maxLeft2))
			}
		} else if maxLeft1 > minRight2 {
			end = p1 - 1
		} else {
			start = p1 + 1
		}
	}
	return 0.0
}
