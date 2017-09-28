package meeting

import (
	"sort"
)

type Interval struct {
	s, e int
}

func CanAttendMeetings(intervals []Interval) bool {
	// sort slice based on start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].s < intervals[j].s
	})
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i].e > intervals[i+1].s {
			return false
		}
	}
	return true
}
