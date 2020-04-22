package set

func subSets(nums []int) [][]int {
	n := len(nums)
	total := 1 << n
	ret := make([][]int, 0, total)
	backtrack(&ret, []int{}, nums, 0)
	return ret
}

// backtrack time complexity # of nodes O(2^N), space complexity O(N)
func backtrack(ret *[][]int, store, nums []int, pos int) {
	*ret = append((*ret), store)
	for i := pos; i < len(nums); i++ {
		ns := len(store)
		store = append(store, nums[i])
		backtrack(ret, store, nums, i+1)
		store = store[:ns]
	}
}

func helper(res *[][]int, nums, temp []int, start, target int) {
	if len(temp) == target {
		bak := make([]int, target)
		copy(bak, temp)
		*res = append(*res, bak)
	} else {
		for i := start; i < len(nums); i++ {
			n := len(temp)
			temp = append(temp, nums[i])
			helper(res, nums, temp, i+1, target)
			temp = temp[:n]
		}
	}
}

func SubSets2(nums []int) [][]int {
	// bit manupulation
	n := len(nums)
	total := 1 << uint(n)
	ret := make([][]int, 0, total)
	for i := 0; i < total; i++ {
		var cur []int
		for j := 0; j < n; j++ {
			if (i>>uint(j))&1 != 0 {
				cur = append(cur, nums[j])
			}
		}
		ret = append(ret, cur)
	}
	return ret
}
