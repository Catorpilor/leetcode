package bitrange

func rangeBitwiseAnd(m, n int) int {
    return commonPrefix(m, n)
}

// commonPrefix time complexity O(1), space complexity O(1)
func commonPrefix(m, n int) int {
    var shift int
    for m < n {
        m >>= 1
        n >>= 1
        shift++
    }
    return m << shift
}
