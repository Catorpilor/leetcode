package pb

import "fmt"

func possibleBipartition(n int, dis [][]int) bool {
	return useHashmap(n, dis)
}

func useHashmap(n int, dis [][]int) bool {
	g1, g2 := make(map[int]bool, n), make(map[int]bool, n)
	for _, p := range dis {
		fmt.Printf("g1: %v and g2: %v\n", g1, g2)
		p0, p1 := p[0], p[1]
		if g1[p0] == true && g1[p1] == true || g2[p1] == true && g2[p0] == true {
			return false
		}
		if !g2[p0] {
			g1[p0] = true
		}
		if !g1[p1] {
			g2[p1] = true
		}
	}
	return true
}
