package intersection

func Intersection(nums1, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	hm := make(map[int]int)
	for _, v := range nums1 {
		if _, ok := hm[v]; !ok {
			hm[v] = 1
		} else {
			hm[v] += 1
		}
	}
	var ret []int
	for _, k := range nums2 {
		if _, ok := hm[k]; ok {
			ret = append(ret, k)
			if hm[k] == 1 {
				delete(hm, k)
			} else {
				hm[k] -= 1
			}
		}
	}
	return ret
}
