package nqueue

import (
	"fmt"
	"strings"
)

type point struct {
	r, c int
}

func solveNQueues(n int) [][]string {
	var vps [][]point

	permute(&vps, []point{}, 0, n)
	fmt.Printf("final queue positions are: %v\n", vps)
	if len(vps) < 1 {
		return [][]string{}
	}

	// construct the final string
	res := make([][]string, len(vps))
	for i := range res {
		res[i] = make([]string, n)
		for j := 0; j < n; j++ {
			dots := strings.Repeat(".", n)
			out := []rune(dots)
			pos := vps[i][j]
			out[pos.c] = 'Q'
			res[i][j] = string(out)
		}

	}
	return res
}

func isValid(qps []point, crow, col int) bool {
	for _, p := range qps {
		fmt.Printf("queue pos: %v, crow: %d, col: %d\n", p, crow, col)
		if p.c == col || crow+col == p.r+p.c || crow-col == p.r-p.c {
			return false
		}
	}
	fmt.Printf("qps :%v and valid crow: %d, col: %d\n", qps, crow, col)
	return true
}

func permute(vmaz *[][]point, qps []point, nrow, nqueue int) {
	if len(qps) == nqueue {
		tmp := make([]point, nqueue)
		copy(tmp, qps)
		*vmaz = append(*vmaz, tmp)
		return
	}
	for j := 0; j < nqueue; j++ {
		if isValid(qps, nrow, j) {
			pl := len(qps)
			qps = append(qps, point{nrow, j})
			permute(vmaz, qps, nrow+1, nqueue)
			qps = qps[:pl]
		}
	}
}
