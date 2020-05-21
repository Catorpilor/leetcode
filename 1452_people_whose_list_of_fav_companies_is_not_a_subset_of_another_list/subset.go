package subset

import "sort"

func peopleIndexes(favCompanies [][]string) []int {
    return useHashmap(favCompanies)
}

// useHashmap time complexity O(N^2*M), space complexity O(N*500)
func useHashmap(favCompanies [][]string) []int {
    n := len(favCompanies)
    collections := make([]map[string]bool, n)
    // O(MN)
    for i := range favCompanies {
        if collections[i] == nil {
            collections[i] = make(map[string]bool, 500)
        }
        for _, comp := range favCompanies[i] {
            collections[i][comp] = true
        }
    }
    ans := make([]int, 0, n)
    for i := 0; i < n; i++ {
        if isSubset(i, favCompanies[i], collections) {
            continue
        }
        ans = append(ans, i)
    }
    return ans
}

//
func isSubset(idx int, companies []string, collection []map[string]bool) bool {
    var count int
    for i := range collection { // O(N)
        if i == idx || len(companies) > len(collection[i]) {
            count++
            continue
        }
        for _, comp := range companies { //O(M)
            if !collection[i][comp] { // O(1)
                count++
                break
            }
        }
    }
    return count != len(collection)
}

// useUnionFind time complexity O(MN+N^2*M), space complexity O(MN)
func useUnionFind(companies [][]string) []int {
    n := len(companies)
    collections := make([]map[string]bool, n)
    // O(MN)
    for i := range companies {
        if collections[i] == nil {
            collections[i] = make(map[string]bool)
        }
        for _, comp := range companies[i] {
            collections[i][comp] = true
        }
    }
    // union find with path compression
    uf := make([]int, n)
    for i := range uf {
        uf[i] = i
    }
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            rooti, rootj := find(uf, i), find(uf, j)
            if rooti == rootj {
                continue
            } else if contains(collections[i], collections[j]) {
                uf[j] = i
            } else if contains(collections[j], collections[i]) {
                uf[i] = j
            }
        }
    }
    res := make([]int, 0, n)
    set := make(map[int]bool, n)
    for i := range uf {
        set[find(uf, i)] = true
    }
    for k := range set {
        res = append(res, k)
    }
    sort.Ints(res)
    return res
}

func contains(s1, s2 map[string]bool) bool {
    n1, n2 := len(s1), len(s2)
    if n1 < n2 {
        return false
    }
    for k := range s2 {
        if !s1[k] {
            return false
        }
    }
    return true
}

func find(uf []int, i int) int {
    for uf[i] != i {
        uf[i] = uf[uf[i]]
        i = uf[i]
    }
    return i
}
