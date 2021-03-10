package roman

import "bytes"

func intToRoman(num int) string {
	return useIter(num)
}

func useIter(num int) string {
	var sb bytes.Buffer
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; num != 0; i++ {
		for num >= vals[i] {
			num -= vals[i]
			sb.WriteString(symbols[i])
		}
	}
	return sb.String()
}
