package ptw

import "github.com/catorpilor/LeetCode/utils"

func predictTheWinner(nums []int) bool {
	return helper(nums, 0, len(nums)-1) >= 0
}

// helper implements maxmin algorithm
// each turn player has two choices nums[s] or nums[e]
// player 1 wins ==> sum of #1 picks >= sum of #2 picks
// aka. sum of #1 picks - sum of #2 picks >= 0
// so when player 1 picks we add it to the total sum, then it's player 2's turn we just subtract it form the total sum
// eg. player 1 picks nums[s] , player 2's picks is helper(nums, s+1,e)
// time complexity is O(2^n) each time we have two choices.
// space complexity is O(n), the recursion tree can only be n length deep.
func helper(nums []int, s, e int) int {
	if s == e {
		return nums[e]
	}
	return utils.Max(nums[s]-helper(nums, s+1, e), nums[e]-helper(nums, s, e-1))
}
