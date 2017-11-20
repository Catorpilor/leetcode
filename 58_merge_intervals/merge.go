package merge

import "sort"

type interval struct {
	start, end int
}

func Merge(intervals []interval) []interval {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}
	// sort intervals based on the start val in each bucket
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})
	ret := make([]interval, 0, n)
	ret = append(ret, intervals[0])
	m := len(ret)
	for i := 1; i < n; i++ {
		// always compare with the last bucket in ret
		m = len(ret)
		if ret[m-1].end < intervals[i].end {

			if ret[m-1].end >= intervals[i].start {
				ret[m-1].end = intervals[i].end
			} else {
				ret = append(ret, intervals[i])
			}
		}
	}
	return ret
}
