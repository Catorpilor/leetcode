package diagonal

import "github.com/catorpilor/leetcode/utils"

func diagonalTraverse(nums [][]int) []int {
    return useHashmap(nums)
}

// func useBruteForce(nums [][]int) []int {
//     m := len(nums)
//     if m < 1 {
//         return nil
//     }
//     if m == 1 {
//         return nums[0]
//     }
//     n := len(nums[0])
//     k := 0
//     upper := m - 1 + n - 1
//     var ans []int
//     for k <= upper {
//         for i, j := k, 0; i+j == k && i >= 0 && j <= k; i, j = i-1, j+1 {
//             if m+len(nums[i])-2 > upper {
//                 upper = m + len(nums[i]) - 2
//             }
//             if j >= len(nums[i]) {
//                 continue
//             }
//             ans = append(ans, nums[i][j])
//         }
//         k++
//     }
//     return ans
// }

// time complexity O(N) N means the number of values, space complexity O(N)
func useHashmap(nums [][]int) []int {
    m := len(nums)
    if m < 1 {
        return nil
    }
    if m == 1 {
        return nums[0]
    }
    set := make(map[int][]int)
    maxK := 0
    for i := range nums {
        for j := range nums[i] {
            maxK = utils.Max(maxK, i+j)
            set[i+j] = append(set[i+j], nums[i][j])
        }
    }
    var ans []int
    for k := 0; k <= maxK; k++ {
        tmp := set[k]
        for l, r := 0, len(tmp)-1; l < r; l, r = l+1, r-1 {
            tmp[l], tmp[r] = tmp[r], tmp[l]
        }
        ans = append(ans, tmp...)
    }
    return ans
}
