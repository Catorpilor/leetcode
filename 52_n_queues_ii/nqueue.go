package nqueue

import "fmt"

func totalNQueues(n int) int {
	var vps [][]point
	permute(&vps, []point{}, 0, n)
	return len(vps)
}

type point struct {
	r, c int
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
