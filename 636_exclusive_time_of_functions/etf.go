package etf

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	return useStack(n, logs)
}

// useStack time complexity O(M) M is the length of logs, space complexity is O(M)
func useStack(n int, logs []string) []int {
	m := len(logs)
	res := make([]int, n)
	st := make([]int, 0, 2000)
	id, _, prev := helper(logs[0])
	st = append(st, id)
	for i := 1; i < m; i++ {
		cid, status, cts := helper(logs[i])
		ns := len(st)
		if status == "start" {
			if ns > 0 {
				// update the top of the stack's time
				res[st[ns-1]] += cts - prev
			}
			st = append(st, cid)
			prev = cts
		} else {
			res[st[ns-1]] += cts - prev + 1
			st = st[:ns-1]
			prev = cts + 1
		}
	}
	return res
}

// helper parse the log in `id:status:ts` format, and returns the fields
func helper(log string) (id int, status string, ts int) {
	arr := strings.FieldsFunc(log, func(r rune) bool {
		return r == ':'
	})
	id, _ = strconv.Atoi(arr[0])
	ts, _ = strconv.Atoi(arr[2])
	status = arr[1]
	return
}
