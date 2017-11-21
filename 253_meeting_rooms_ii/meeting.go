package meeting

import (
	"fmt"
	"math"
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

type interval struct {
	Start, End int
}

func MinRooms(intervals []interval) int {
	n := len(intervals)
	if n <= 1 {
		return n
	}
	// sort intervals based on their Start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	bfound := false
	ret := 1
	hash := make(map[int]bool)
	if intervals[1].Start < intervals[0].End {
		hash[intervals[0].End], hash[intervals[1].End], ret = true, true, 2
	}
	for i := 2; i < n; i++ {
		bfound = false
		if intervals[i].Start < intervals[i-1].End {
			// overlapping
			min := math.MaxInt32
			for k, _ := range hash {
				if intervals[i].Start >= k {
					bfound = true
					if k < min {
						min = k
					}
				}
			}
			// no rooms found
			if !bfound {
				ret += 1
				// hash[intervals[i].End] = true
			} else {
				delete(hash, min)
				// hash[intervals[i].End] = true
			}
		} else {
			delete(hash, intervals[i-1].End)
			// hash[intervals[i].End] = true
		}
		hash[intervals[i].End] = true
	}
	// fmt.Println(hash)

	return ret
}

func MinRooms2(intervals []interval) int {
	n := len(intervals)
	if n <= 1 {
		return n
	}
	// sort intervals based on their Start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	rooms := make([]int, 0, n)
	if intervals[0].End > intervals[1].Start {
		rooms = append(rooms, intervals[0].End, intervals[1].End)
	} else {
		rooms = append(rooms, intervals[1].End)
	}
	bfound := false
	for i := 2; i < n; i++ {
		bfound = false
		// if intervals[i].Start < intervals[i-1].End {
		for j := range rooms {
			if intervals[i].Start >= rooms[j] {
				bfound = true
				rooms[j] = intervals[i].End
				break
			}
		}
		if !bfound {
			rooms = append(rooms, intervals[i].End)
		}
		// }
	}
	return len(rooms)
}

func MinRooms3(intervals []interval) int {
	n := len(intervals)
	if n <= 1 {
		return n
	}
	times := make([]int, n*2)
	for i := range intervals {
		times[2*i] = intervals[i].Start
		times[2*i+1] = -intervals[i].End
	}
	sort.Slice(times, func(i, j int) bool {
		if utils.Abs(times[i]) == utils.Abs(times[j]) {
			return i < j
		}
		return utils.Abs(times[i]) < utils.Abs(times[j])
	})
	var cur, ret int
	for i := range times {
		if times[i] >= 0 {
			cur += 1
			ret = utils.Max(ret, cur)
		} else {
			cur -= 1
		}
	}
	fmt.Println(times)
	return ret
}

func MinRooms4(intervals []interval) int {
	n := len(intervals)
	if n <= 1 {
		return n
	}
	start, end := make([]int, n), make([]int, n)
	for i := range intervals {
		start[i], end[i] = intervals[i].Start, intervals[i].End
	}
	sort.Slice(start, func(i, j int) bool {
		return start[i] < start[j]
	})
	sort.Slice(end, func(i, j int) bool {
		return end[i] < end[j]
	})
	var rooms, endIdx int
	for i := range start {
		if start[i] < end[endIdx] {
			rooms += 1
		} else {
			endIdx += 1
		}
	}
	return rooms
}
