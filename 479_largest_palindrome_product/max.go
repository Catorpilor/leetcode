package max

import (
	"strconv"
	"math"
)


func LargestPalindrome(n int) int{
	if n <= 1 {
		return 9
	}
	maxNum := int64(math.Pow10(n) - 1)
	minNum := int64(math.Pow10(n-1))
	maxProduct := maxNum * maxNum
	firstHalf := maxProduct / int64(math.Pow10(n))
	for {
		candidate := palindrome(firstHalf)
		if candidate > maxProduct {
			firstHalf -= 1
			continue
		}
		for i:=maxNum; i>=minNum; i--{
			if candidate / i > maxNum {
				break
			}
			if candidate % i == 0 {
				return int(candidate % 1337)
			}
		}
		firstHalf -= 1
	}
}

func palindrome(fh int64) int64{
	sfh := strconv.FormatInt(fh, 10)
	pli := "" + sfh
	runes := []rune(sfh)
	// reverse sfh
	for i,j := 0,len(runes)-1;i<j;i,j=i+1,j-1 {
		runes[i],runes[j] = runes[j],runes[i]
	}
	pli += string(runes)
	p64, _ := strconv.ParseInt(pli, 10, 64)
	return p64
}