package matrix

import (
	"fmt"
	"sort"
)

func diagonalSort(matrix [][]int) [][]int {
	return useBruteForce(matrix)
}

// useBruteForce time complexity O((M+N)*K*log(K)) where K is the average length of the diagonal length. space complexity O(MN)
func useBruteForce(matrix [][]int) [][]int {
	m := len(matrix)
	if m < 1 {
		return matrix
	}
	n := len(matrix[0])
	if n < 1 {
		return matrix
	}
	for i := m - 2; i >= 0; i-- {
		tmp := make([]int, 0, n<<1)
		for x, y := i, 0; x < m && y < n; x, y = x+1, y+1 {
			tmp = append(tmp, matrix[x][y])
		}
		sort.Ints(tmp)
		fmt.Printf("at row %d, after sort got tmp:%v\n", i, tmp)
		x, y := i, 0
		for _, num := range tmp {
			matrix[x][y] = num
			fmt.Printf("change matrix(%d,%d) to num:%d\n", x, y, num)
			x, y = x+1, y+1
		}
	}
	for j := 1; j < n-1; j++ {
		tmp := make([]int, 0, n<<1)
		for x, y := 0, j; x < m && y < n; x, y = x+1, y+1 {
			tmp = append(tmp, matrix[x][y])
		}
		sort.Ints(tmp)
		fmt.Printf("at col %d, after sort got tmp:%v\n", j, tmp)
		x, y := 0, j
		for _, num := range tmp {
			matrix[x][y] = num
			fmt.Printf("change matrix(%d,%d) to num:%d\n", x, y, num)
			x, y = x+1, y+1
		}
	}
	fmt.Println(matrix)
	return matrix
}
