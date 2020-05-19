package subset

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
