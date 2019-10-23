package permute

import "fmt"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	if nums == nil || len(nums) == 0 {
		return res
	}

	hset := make(map[int]int)
	for i := range nums {
		hset[nums[i]]++
	}
	permute(&res, hset, []int{}, len(nums))
	return res

}

func permute(res *[][]int, hset map[int]int, r []int, n int) {
	if n <= 0 {
		fmt.Printf("find one, r is : %v, res: %v\n", r, *res)
		tmp := make([]int, len(r))
		copy(tmp, r)
		*res = append(*res, tmp)
		return
	}
	for k := range hset {
		if hset[k] > 0 {
			hset[k]--
			r = append(r, k)
			permute(res, hset, r, n-1)
			r = r[:len(r)-1]
			hset[k]++
		}
	}
	return
}

func permuteUnique2(nums []int) [][]int {
	answer := make([][]int, 0)
	aux(&answer, 0, nums)
	return answer
}

func aux(answer *[][]int, idx int, nums []int) {
	if idx == len(nums) {
		c := make([]int, len(nums))
		copy(c, nums)
		*answer = append(*answer, c)
		return
	}
	visited := make(map[int]bool, 0)
	for i := idx; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}
		nums[idx], nums[i] = nums[i], nums[idx]
		aux(answer, idx+1, nums)
		nums[i], nums[idx] = nums[idx], nums[i]
		visited[nums[i]] = true
	}
	return
}
