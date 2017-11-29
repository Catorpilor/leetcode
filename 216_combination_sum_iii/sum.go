package sum

import "fmt"

func CombinationSum(k, n int) [][]int {
	// backtracking
	var ret [][]int
	var temp []int
	helper(&ret, &temp, 1, n, k)
	fmt.Println(ret)
	return ret
}

func helper(res *[][]int, temp *[]int, start, n, k int) {
	if k == 0 {
		if n == 0 {
			bak := make([]int, len(*temp))
			copy(bak, *temp)
			*res = append(*res, bak)
		}
	} else {
		for i := start; i <= 9 && n >= i; i++ {
			lt := len(*temp)
			*temp = append(*temp, i)
			helper(res, temp, i+1, n-i, k-1)
			*temp = (*temp)[:lt]
		}
	}
}
