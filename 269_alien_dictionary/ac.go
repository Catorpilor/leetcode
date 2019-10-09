package ac

import "fmt"

func alienOrder(words []string) string {
	var ans string
	n := len(words)
	indegress := make(map[byte]int)
	g := make(map[byte]map[byte]bool)
	// init indegrees
	for i := 0; i < n; i++ {
		for j := 0; j < len(words[i]); j++ {
			indegress[words[i][j]] = 0
		}
	}
	// build graph and update indegress
	for i := 0; i < n-1; i++ {
		curW, nxtW := words[i], words[i+1]
		sn := len(curW)
		if sn > len(nxtW) {
			sn = len(nxtW)
		}
		fmt.Printf("curW: %s, nextW: %s, sn: %d\n", curW, nxtW, sn)
		for j := 0; j < sn; j++ {
			if curW[j] != nxtW[j] {
				fmt.Printf("find one: left: %s, right: %s, j: %d\n", string(curW[j]), string(nxtW[j]), j)
				nodes := g[curW[j]]
				if nodes == nil {
					nodes = make(map[byte]bool)
				}
				if nodes[nxtW[j]] == false {
					nodes[nxtW[j]] = true
					g[curW[j]] = nodes
					indegress[nxtW[j]]++
				}
				break
			}
		}
	}
	fmt.Printf("graph is %v\n, indegress is %v\n", g, indegress)
	for k := range indegress {
		if indegress[k] == 0 {
			ans += string(k)
		}
	}
	fmt.Printf("current ans: %s\n", ans)
	for i := 0; i < len(ans); i++ {
		nodes := g[ans[i]]
		fmt.Printf("current nodes is: %v\n", nodes)
		for k := range nodes {
			indegress[k]--
			if indegress[k] == 0 {
				ans += string(k)
			}
		}
	}
	if len(ans) != len(indegress) {
		return ""
	}
	return ans
}
