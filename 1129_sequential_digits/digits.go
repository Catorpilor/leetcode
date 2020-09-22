package digits

import "sort"

func seqDigits(low, high int) []int {
	return useStraight(low, high)
}

// useStraight time complexity O(N) where N is the max digits, space complexity O(1)
func useStraight(low, high int) []int {
	var lnd, hnd int
	num := low
	for num > 0 {
		num /= 10
		lnd++
	}
	num = high
	for num > 0 {
		num /= 10
		hnd++
	}
	var ans []int
	for i := 1; i < 9; i++ {
		gen(&ans, i, low, high, hnd)
	}
	sort.Ints(ans)
	return ans
}

func gen(ans *[]int, startNum, low, high, hnd int) {
	var idx, num int
	for idx < hnd {
		if startNum+idx < 10 {
			num = num*10 + startNum + idx
			idx++
			if num >= low && num <= high {
				*ans = append((*ans), num)
			}
		} else {
			break
		}
	}
}
