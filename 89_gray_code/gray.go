package gray

import "github.com/catorpilor/leetcode/utils"

// sovled based on top voted discussion
// generate the sequence iteratively. For example, when n=3, we can get the result based on n=2.
// n = 2 sequence is  00, 01, 11, 10
// n = 3, the total numer is 8, the rest 4 is the symmetric mirror of the previous generated seq
// the only difference is the highest bit
// so it becomoe 000, 001, 011, 010, 110, 111,101, 100
// the time complexity is O(2^n)

func grayCode(n int) []int {
	var res []int
	if n < 0 {
		return res
	}
	res = make([]int, 0, utils.Pow(2, n))
	res = append(res, 0)
	for i := 0; i < n; i++ {
		rn := len(res)
		for j := rn - 1; j >= 0; j-- {
			res = append(res, (res[j] | 1<<i))
		}
	}
	return res
}
