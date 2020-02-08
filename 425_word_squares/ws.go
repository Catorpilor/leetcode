package ws

import (
	"strings"

	"github.com/catorpilor/leetcode/utils"
)

func buildWithWords(words []string) *utils.TernarySearchTrie {
	tst := utils.NewTST()
	for _, w := range words {
		tst.Put(w, w)
	}
	return tst
}

func wordSquares(words []string) [][]string {
	nl := len(words)
	if nl < 1 {
		return [][]string{}
	}
	tst := buildWithWords(words)
	n := len(words[0])
	var res [][]string
	for _, w := range words {
		wsq := make([]string, 0, nl)
		wsq = append(wsq, w)
		helper(&res, &wsq, 1, n, tst)
	}
	return res
}

func helper(res *[][]string, wsq *[]string, idx, n int, tst *utils.TernarySearchTrie) {
	if idx == n {
		tmp := make([]string, len(*wsq))
		copy(tmp, *wsq)
		*res = append(*res, tmp)
		return
	}
	var sb strings.Builder
	for _, w := range *wsq {
		sb.WriteByte(w[idx])
	}
	//fmt.Printf("idx:%d, n: %d, wsq: %v, prefix:%s\n", idx, n, *wsq, sb.String())
	vals := tst.WithPrefix(sb.String())
	// fmt.Printf("idx:%d, n: %d, wsq: %v, prefix:%s,vals: %v\n", idx, n, *wsq, sb.String(), vals)
	for _, v := range vals {
		nn := len(*wsq)
		*wsq = append(*wsq, v.(string))
		helper(res, wsq, idx+1, n, tst)
		*wsq = (*wsq)[:nn]
	}
}
