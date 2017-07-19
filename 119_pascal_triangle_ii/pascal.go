package pascal

func GetRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}
	if rowIndex == 0 {
		return []int{1}
	}
	prev := []int{1, 1}
	if rowIndex == 1 {
		return prev
	}
	for i := 2; i <= rowIndex; i++ {
		p := []int{1}
		for j := 0; j < len(prev)-1; j++ {
			p = append(p, prev[j]+prev[j+1])
		}
		p = append(p, 1)
		prev = p
	}
	return prev
}
