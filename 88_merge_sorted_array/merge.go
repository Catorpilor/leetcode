package merge


func Merge(nums1, nums2 []int, m,n int) []int {
	i,j,tar := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j]{
				nums1[tar] = nums1[i]
				i -= 1
		}else {
			nums1[tar] = nums2[j]
			j -= 1
		}
		tar -= 1
	}
	return nums1
}