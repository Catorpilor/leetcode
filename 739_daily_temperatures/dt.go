package dt

func dailyTemps(temps []int) []int {
    return useStack(temps)
}

// useStack time complexity O(N), space complexity O(N)
func useStack(temps []int) []int {
    n := len(temps)
    ans := make([]int, n)
    if n <= 1 {
        return ans
    }
    st := make([]int, 0, n)
    for i := range temps {
        nst := len(st)
        for nst > 0 && temps[st[nst-1]] < temps[i] {
            // pop from stack
            top := st[nst-1]
            ans[top] = i - top
            st = st[:nst-1]
            nst--
        }
        st = append(st, i)
    }
    return ans
}
