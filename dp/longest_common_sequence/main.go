package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/catorpilor/leetcode/utils"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var m, n int
	fmt.Fscan(in, &m, &n)
	arr1, arr2 := make([]int, m), make([]int, n)
	for i := range arr1 {
		fmt.Fscan(in, &arr1[i])
	}
	for j := range arr2 {
		fmt.Fscan(in, &arr2[j])
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if arr1[i-1] == arr2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = utils.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	// ret := helper(dp, arr1, arr2, m, n, dp[m][n])
	// for i := len(ret) - 1; i >= 0; i-- {
	// 	fmt.Printf("%d ", ret[i])
	// }
	var ret [][]int
	temp := make([]int, 0, dp[m][n])
	helperAll(dp, &ret, temp, arr1, arr2, m, n, dp[m][n])
	fmt.Println(ret)

}

func helper(dp [][]int, arr1, arr2 []int, i, j, target int) []int {
	ret := make([]int, 0, target)
	for i != 0 && j != 0 {
		if arr1[i-1] == arr2[j-1] {
			ret = append(ret, arr1[i-1])
			i, j = i-1, j-1
		} else {
			if dp[i][j-1] > dp[i-1][j] {
				j--
			} else {
				i--
			}
		}
	}
	return ret
}

func helperAll(dp [][]int, res *[][]int, temp []int, arr1, arr2 []int, i, j, target int) {
	if target == 0 {
		bak := make([]int, len(temp))
		copy(bak, temp)
		*res = append(*res, bak)
	} else {
		if arr1[i-1] == arr2[j-1] {
			temp = append(temp, arr1[i-1])
			helperAll(dp, res, temp, arr1, arr2, i-1, j-1, target-1)
		} else {
			if dp[i][j-1] >= dp[i-1][j] {
				helperAll(dp, res, temp, arr1, arr2, i, j-1, target)
			}
			if dp[i-1][j] >= dp[i][j-1] {
				helperAll(dp, res, temp, arr1, arr2, i-1, j, target)
			}
		}
	}
}
