package intersection

func Intersection(nums1, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	hm := make(map[int]bool)
	for _, v := range nums1 {
		hm[v] = true
	}
	var ret []int
	for _, k := range nums2 {
		if hm[k] {
			ret = append(ret, k)
			delete(hm, k)
		}
	}
	return ret
}
