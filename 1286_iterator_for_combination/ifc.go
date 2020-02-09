package ifc

import (
	"github.com/catorpilor/leetcode/utils"
)

type CombinationIterator struct {
	reversedStr string
	cl          int
	bmask       int
}

func (c *CombinationIterator) Next() string {
	for c.bmask >= 0 && utils.NumOfOnes(c.bmask) != c.cl {
		c.bmask--
	}
	var res string
	for i := 0; i < len(c.reversedStr); i++ {
		if ((c.bmask & (1 << i)) >> i) != 0 {
			res = string(c.reversedStr[i]) + res
		}
	}
	//fmt.Printf("bMask: %d and res: %s\n", c.bmask, res)
	c.bmask--
	return res
}

func (c *CombinationIterator) HasNext() bool {
	for c.bmask >= 0 && utils.NumOfOnes(c.bmask) != c.cl {
		c.bmask--
	}
	if c.bmask < 0 {
		return false
	}
	return true
}

func NewCombineIterator(s string, combineLength int) *CombinationIterator {
	sb := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		sb[i], sb[j] = sb[j], sb[i]
	}
	return &CombinationIterator{
		reversedStr: string(sb),
		cl:          combineLength,
		bmask:       (1 << len(s)) - 1,
	}
}
