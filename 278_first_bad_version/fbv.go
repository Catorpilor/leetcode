package fbv

func firstBadVersion(vers []int) int {
    return useBinarySearch(vers)
}

// useBinarySearch time complexity O(lgN), space compelxity O(1)
func useBinarySearch(vers []int) int {
    l, r := 0, len(vers)
    ans := -1
    for l <= r {
        mid := l + (r-l)/2
        f := isBadVersion(vers[mid])
        if !f {
            l = mid + 1
        } else {
            ans = mid
            r = mid - 1
        }
    }
    return ans
}

func isBadVersion(n int) bool {
    return n == 0
}
