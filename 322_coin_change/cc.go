package cc

import (
	"fmt"
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

const MAX_INT = 1<<31 - 1

func coinChange(coins []int, amount int) int {
	return useBucket(coins, amount)
}

// useBucket time complexity O(MN) space complxity O(N) where M is the length of coins, N is the number of amount.
func useBucket(coins []int, amount int) int {
	if len(coins) < 1 {
		return -1
	}
	sort.Slice(coins, func(i, j int) bool { return coins[i] <= coins[j] })
	ta := make([]int, amount+1) // returns the minimum count for coin in coins to sum up as amount
	ta[0] = 0

	// assume coins are sorted in increasing order
	for j := 1; j <= amount; j++ {
		if j < coins[0] {
			ta[j] = MAX_INT
		} else {
			if ta[j%coins[0]] == MAX_INT {
				ta[j] = MAX_INT
			} else {
				ta[j] = j / coins[0]
			}
		}
	}
	fmt.Println(ta[amount])
	for i := 1; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i] {
				ta[j] = utils.Min(ta[j], ta[j-coins[i]]+1)
			}
		}
	}
	if ta[amount] == MAX_INT {
		return -1
	}
	return ta[amount]
}

func CoinChange2(coins []int, amount int) int {
	// ta stores the minimal counts of coins to sum up i
	ta := make([]int, amount+1)
	ta[0] = 0

	for i := 1; i <= amount; i++ {
		// set ta[i] to amount (MAX)
		ta[i] = amount + 1
		for _, c := range coins {
			if i >= c && ta[i-c]+1 < ta[i] {
				ta[i] = ta[i-c] + 1
			}
		}
	}
	if ta[amount] > amount {
		return -1
	}
	return ta[amount]
}
