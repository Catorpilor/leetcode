package sum

type NumArray struct {
	sums []int
}

func NewNumArray(nums []int) *NumArray {
	l := len(nums)
	if l == 0 {
		return nil
	}
	prev, cur := 0, 0
	sums := make([]int, l)
	for i := 0; i < l; i++ {
		cur = prev + nums[i]
		sums[i] = cur
		prev = cur
	}
	return &NumArray{sums: sums}
}

func (n *NumArray) SumRange(i, j int) int {
	if i == 0 {
		return n.sums[j]
	}
	return n.sums[j] - n.sums[i-1]
}
