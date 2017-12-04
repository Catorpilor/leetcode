package cal

import "fmt"

type Calendar struct {
	starts, ends []int
}

func Constructor() Calendar {
	return Calendar{starts: make([]int, 0, 1000), ends: make([]int, 0, 1000)}
}

func (c *Calendar) Book(start, end int) bool {
	if start > end {
		return false
	}
	if len(c.starts) == 0 {
		c.starts, c.ends = append(c.starts, start), append(c.ends, end)
		return true
	}
	pos1, pos2 := lower_bound(c.starts, 0, len(c.starts), start),
		lower_bound(c.ends, 0, len(c.ends), end)
	if pos1 == -1 || pos2 == -1 || pos1 != pos2 {
		return false
	}
	if pos1 == 0 {
		if end > c.starts[pos1] {
			return false
		}
	} else if pos1 == len(c.starts) {
		if start < c.ends[pos1-1] {
			return false
		}
	} else {
		if start < c.ends[pos1-1] || end > c.starts[pos1] {
			return false
		}
	}
	c.starts = append(c.starts, 0)
	copy(c.starts[pos1+1:], c.starts[pos1:])
	c.starts[pos1] = start
	c.ends = append(c.ends, 0)
	copy(c.ends[pos2+1:], c.ends[pos2:])
	c.ends[pos2] = end
	// c.starts = append(append(c.starts[:pos1], start), c.starts[pos1:]...)
	// c.ends = append(append(c.ends[:pos2], end), c.ends[pos2:]...)
	return true
}

func (c *Calendar) Status() {
	fmt.Printf("starts: %v, ends: %v\n", c.starts, c.ends)
}

func lower_bound(nums []int, low, hi, target int) int {
	for low < hi {
		mid := low + (hi-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] == target {
			return -1
		} else {
			hi = mid
		}
	}
	return low
}

type event struct {
	start, end int
}

type MyCalendar struct {
	events []event
}

func NewCalender() MyCalendar {
	return MyCalendar{events: make([]event, 0, 1000)}
}

func (c *MyCalendar) Book(start, end int) bool {
	for _, e := range c.events {
		if !(e.end <= start || e.start >= end) {
			return false
		}
	}
	c.events = append(c.events, event{start: start, end: end})
	return true
}
