package calls

type RecentCalls struct {
	calls []int
}

func Constructor() *RecentCalls {
	return &RecentCalls{}
}

func (rc *RecentCalls) Ping(t int) int {
	return -1
}
