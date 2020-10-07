package calls

import "github.com/catorpilor/leetcode/utils"

type RecentCalls struct {
	calls []int
}

func Constructor() *RecentCalls {
	return &RecentCalls{}
}

// Ping time complexity O(lgN), space complexity O(1)
func (rc *RecentCalls) Ping(t int) int {
	rc.calls = append(rc.calls, t)
	l := utils.LowerBound(rc.calls, t-3000)
	return len(rc.calls) - l
}
