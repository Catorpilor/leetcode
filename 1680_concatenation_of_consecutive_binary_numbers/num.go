package num

import (
	"strconv"
)

func concatenateBinary(n int) int {
	return useMath(n)
}

const (
	mod = int64(1e9 + 7)
)

func useMath(n int) int {
	var ans int64
	for i := 1; i <= n; i++ {
		bi := strconv.FormatInt(int64(i), 2)
		ans = ((ans << len(bi)) + int64(i)) % mod
	}
	return int(ans)
}
