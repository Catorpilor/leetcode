package mlrs

import (
	"reflect"

	"github.com/catorpilor/leetcode/utils"
)

func MaxLength(nums1, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return 0
	}
	m, n := n1, n2
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	var ret int
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				ret = utils.Max(ret, dp[i][j])
			}
		}
	}
	return ret
}

func MaxLength2(nums1, nums2 []int) int {
	// rolling hash
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return 0
	}
	low, hi := 0, utils.Min(n1, n2)+1
	for low < hi {
		mid := (hi-low)/2 + low
		if check(mid, nums1, nums2) {
			low = mid + 1
		} else {
			hi = mid
		}

	}
	if low > 0 {
		return low - 1
	}
	return 0
}

const base = 16777619

func check(length int, nums1, nums2 []int) bool {
	hash := make(map[uint32][]int)
	for i, v := range rollingHash(nums1, length) {
		if _, ok := hash[v]; ok {
			hash[v] = append(hash[v], i)
		} else {
			hash[v] = []int{i}
		}
	}
	for j, v := range rollingHash(nums2, length) {
		for _, i := range hash[v] {
			if reflect.DeepEqual(nums1[i:i+length], nums2[j:j+length]) {
				return true
			}
		}
	}
	return false
}

func rollingHash(nums []int, l int) []uint32 {
	if l == 0 {
		return nil
	}
	var multi uint32 = 1
	for i := 0; i < l-1; i++ {
		multi *= base
	}

	ret := make([]uint32, len(nums)-l+1)
	ret[0] = hash(nums[:l])
	for i := 1; i < len(nums)-l+1; i++ {
		ret[i] = (ret[i-1]-multi*uint32(nums[i-1]))*base + uint32(nums[i+l-1])
	}
	return ret
}

func hash(s []int) uint32 {
	var h uint32
	for _, v := range s {
		h = h*base + uint32(v)
	}
	return h
}
