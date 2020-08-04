package winner

import "github.com/catorpilor/leetcode/utils"

func winnerOfGame(arr []int, k int) int {
	return useTwoPassWithStack(arr, k)
}

// useStack time complexity O(N), space complexity O(N)
func useTwoPassWithStack(arr []int, k int) int {
	n := len(arr)
	posOfPeeks := make([]int, 0, n)
	var maxV int
	for i := 0; i < n; i++ {
		if i > 0 && arr[i] < arr[i-1] {
			continue
		}
		if i < n-1 && arr[i] < arr[i+1] {
			continue
		}
		posOfPeeks = append(posOfPeeks, i)
		maxV = utils.Max(maxV, arr[i])
	}
	// len(posOfPeeks) means the arr is strictly increasing or decreasing
	if len(posOfPeeks) == 1 || n < k {
		return maxV
	}
	st := make([]int, 0, n)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		nst := len(st)
		if nst > 0 {
			top := st[nst-1]
			if arr[top] > arr[i] {
				count[top]++
				if count[top] >= k {
					return arr[top]
				}
				continue
			} else {
				st = st[:nst-1]
				nst--
				count[i]++
			}
		}
		st = append(st, i)
	}
	return maxV
}

// useOnePassWithStack time complexity O(N), space complexity O(N)
func useOnePassWithStack(arr []int, k int) int {
	n := len(arr)
	st := make([]int, 0, n)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		nst := len(st)
		if nst > 0 {
			top := st[nst-1]
			if arr[top] > arr[i] {
				count[top]++
				if count[top] >= k {
					return arr[top]
				}
				continue
			} else {
				st = st[:nst-1]
				nst--
				count[i]++
				if count[i] >= k {
					return arr[i]
				}
			}
		}
		st = append(st, i)
	}
	return arr[st[0]]
}

// useOnePass time complexity O(N), space complexity O(1)
func useOnePass(arr []int, k int) int {
	n := len(arr)
	cur := arr[0]
	wins := 0
	for i := 1; i < n; i++ {
		if arr[i] > cur {
			cur = arr[i]
			wins = 0
		}
		wins++
		if wins >= k {
			break
		}
	}
	return cur
}
