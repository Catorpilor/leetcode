package course

import "fmt"

func minSems(n int, deps [][]int, k int) int {
	return useTopSort(n, deps, k)
}

// useTopSort
func useTopSort(n int, dependencies [][]int, k int) int {
	nd := len(dependencies)
	if nd == 0 {
		return n/k + n%k
	}
	adj := make([][]int, n+1)
	indgree := make([]int, n+1)
	// we need to calculagte the out degrees
	for _, dep := range dependencies {
		p, q := dep[0], dep[1]
		adj[p] = append(adj[p], q)
		indgree[q]++
	}
	visited := make([]bool, n+1)
	var ans, count int
	// find ingree == 0
	for {
		starts := make([]int, 0, n+1)
		for i := 1; i <= n; i++ {
			if indgree[i] == 0 && !visited[i] {
				starts = append(starts, i)
				visited[i] = true
				count++
			}
		}
		fmt.Printf("pick starts: %v, count: %d\n", starts, count)
		// fmt.Println(starts)
		// this logic is wrong.....
		ns := len(starts)
		if ns <= k {
			ans++
		} else {
			ans += ns/k + ns%k
		}
		for i := range starts {
			child := adj[starts[i]]
			for j := range child {
				indgree[child[j]]--
			}
		}
		if count == n {
			break
		}
	}
	return ans
}
