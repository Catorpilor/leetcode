package task

import (
	"github.com/catorpilor/leetcode/utils"
)

func leastInterval(tasks []byte, n int) int {
	return useGreedy(tasks, n)
}

// useGreedy time complexity O(N), space complexity O(1)
func useGreedy(tasks []byte, n int) int {
	hm := make([]int, 26)
	var id, maxF, numOfMaxF int
	for _, b := range tasks {
		id = int(b - 'A')
		hm[id]++
		if hm[id] > maxF {
			maxF = hm[id]
			numOfMaxF = 1
		} else if hm[id] == maxF {
			numOfMaxF++
		}
	}
	// formular (maxF - 1) * (n+1) + numOfMaxf
	// for example with input tasks ['A','A','A','B','B','C'] and n = 3
	// so 'A' has the max frequency illustrated as follows
	// A***A***A
	// then the less frequency one is 'B', we insert B to the slot
	// AB**AB**A
	// then insert 'C'
	// ABC*AB**A so the leastInterval is 9
	// maxF - 1  means how many loops (A***)(A***)
	// n + 1 means the length of each loop A*** = 4 = n + 1
	// numOfMaxF means additional slots needed (A) = 1
	return utils.Max(len(tasks), (maxF-1)*(n+1)+numOfMaxF)
}
