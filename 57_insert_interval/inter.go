package inter

import "github.com/catorpilor/leetcode/utils"

type Interval struct {
	Start, End int
}

func Insert(intervals []Interval, newInterval Interval) []Interval {
	var ret []Interval
	var idx int
	n := len(intervals)
	// add all the intervals ending befor newInterval starts
	for idx < n && intervals[idx].End < newInterval.Start {
		ret = append(ret, intervals[idx])
		idx++
	}
	// mrege the overlapped ones
	for idx < n && intervals[idx].Start <= newInterval.End {
		newInterval = Interval{
			Start: utils.Min(intervals[idx].Start, newInterval.Start),
			End:   utils.Max(intervals[idx].End, newInterval.End),
		}
		idx++
	}
	ret = append(ret, newInterval)
	if idx < n {
		// add teh rest
		ret = append(ret, intervals[idx:]...)
	}
	return ret
}
