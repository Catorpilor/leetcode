package ladder

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func Ladders(start, end string, ws []string) [][]string {
	var ret [][]string
	temp := []string{start}
	dict := make(map[string]bool) // maps ws for quick check
	for i := range ws {
		dict[ws[i]] = true
	}
	ns := make(map[string][]string)  // store neighbors for key
	distance := make(map[string]int) // distance from the start
	dict[start] = true
	bfs(start, end, dict, ns, distance)
	helper(&ret, dict, &temp, ns, distance, start, end)
	return ret
}

func helper(res *[][]string, dict map[string]bool, temp *[]string, ns map[string][]string, distance map[string]int, pivotS, es string) {
	// n := len(*temp)
	// *temp = append(*temp, pivotS)
	if pivotS == es {
		bak := make([]string, len(*temp))
		copy(bak, *temp)
		*res = append(*res, bak)
	} else {
		for _, s := range ns[pivotS] {
			if distance[s] == distance[pivotS]+1 {
				n := len(*temp)
				*temp = append(*temp, s)
				helper(res, dict, temp, ns, distance, s, es)
				*temp = (*temp)[:n]
			}
		}
	}
	// *temp = (*temp)[:n]
}

func bfs(ss, es string, dict map[string]bool, hn map[string][]string, distance map[string]int) {
	q := utils.NewQueue()
	q.Enroll(ss)
	distance[ss] = 0
	var found bool
	for !q.IsEmpty() {
		cnt := q.Size()
		for i := 0; i < cnt; i++ {
			cur := q.Pull().(string)
			curDis := distance[cur]
			ns := neighbors(cur, dict)
			hn[cur] = append(hn[cur], ns...)
			for _, s := range ns {
				if _, ok := distance[s]; !ok {
					distance[s] = curDis + 1
					if s == es {
						found = true
					} else {
						q.Enroll(s)
					}
				}
			}
			if found {
				break
			}
		}
	}
}

func neighbors(s string, dict map[string]bool) []string {
	var ret []string
	bs := []byte(s)
	for c := 'a'; c <= 'z'; c++ {
		for i, cc := range bs {
			if c == rune(cc) {
				continue
			}
			oldC := bs[i]
			bs[i] = byte(c)
			if dict[string(bs)] {
				ret = append(ret, string(bs))
			}
			bs[i] = oldC
		}
	}
	return ret
}

func bfsPath(beg, end string, list []string) [][]string {
	n := len(list)
	if n < 1 {
		return [][]string{}
	}
	words := make(map[string]bool, n)
	for _, str := range list {
		words[str] = true
	}
	if !words[end] {
		return [][]string{}
	}
	var res [][]string
	q := utils.NewQueue()
	visited := make(map[string]bool, n)
	q.Enroll([]string{beg})
	level, minLevel := 1, math.MaxInt32
	for !q.IsEmpty() {
		op := q.Pull().([]string)
		nop := len(op)
		if nop > level {
			// new level
			for k := range visited {
				delete(words, k)
				delete(visited, k)
			}
			if nop > minLevel {
				break
			} else {
				level = nop
			}
		}
		last := op[nop-1]
		// find next words in the list
		for i := 0; i < len(last); i++ {
			nw := []byte(last)
			for c := 'a'; c <= 'z'; c++ {
				nw[i] = byte(c)
				if words[string(nw)] {
					newPath := make([]string, nop+1)
					copy(newPath, op)
					newPath[nop] = string(nw)
					visited[string(nw)] = true
					if string(nw) == end {
						res = append(res, newPath)
						minLevel = level
					} else {
						q.Enroll(newPath)
					}
				}
			}
		}
	}

	return res
}
