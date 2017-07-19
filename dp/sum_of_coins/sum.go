package sum

const maxInt = int(^uint(0) >> 1)

func NumberOfCoins(coins []int, sum int) int {
	if sum <= 0 {
		return 0
	}
	sums := make([]int, sum+1)
	for i := 1; i <= sum; i++ {
		sums[i] = maxInt
	}

	for _, c := range coins {
		for j := 0; j < len(sums); j++ {
			if c+j <= sum && sums[j]+1 < sums[c+j] {
				sums[c+j] = sums[j] + 1
			}
		}
	}
	return sums[sum]
}
