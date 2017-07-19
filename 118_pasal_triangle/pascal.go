package pascal

func Generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}
	ret := [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}}
	if numRows < 4 {
		return ret[:numRows]
	}

	for i := len(ret); i < numRows; i++ {
		t := ret[i-1] // last row in the triangle
		p := []int{1}
		for j := 0; j < len(t)-1; j++ {
			p = append(p, t[j]+t[j+1])
		}
		p = append(p, 1)
		ret = append(ret, p)
	}
	return ret
}
