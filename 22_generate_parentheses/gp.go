package gp

func generateParenthesis(n int) []string {
	var res []string
	permute(&res, "", 0, 0, n)
	return res
}

// use backtracking
// bid stores the current generated string
// leftParenthesis the # of leftParenthesis  (
// rightParenthesis the # of rightParenthesis  )
func permute(res *[]string, bid string, leftParenthesis, rightParenthesis, n int) {
	if len(bid) == n*2 {
		*res = append(*res, bid)
		return
	}
	if leftParenthesis < n {
		permute(res, bid+"(", leftParenthesis+1, rightParenthesis, n)
	}
	if rightParenthesis < leftParenthesis {
		permute(res, bid+")", leftParenthesis, rightParenthesis+1, n)
	}
}
