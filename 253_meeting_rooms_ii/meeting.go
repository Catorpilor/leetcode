package meeting

import (
	"fmt"
	"math"
	"sort"
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
	fmt.Println(rooms)
	return len(rooms)
}
