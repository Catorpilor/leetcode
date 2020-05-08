package check

func isStraight(coors [][]int) bool {
    return useFormular(coors)
}

type f func(int) int

// useFormular time complexity O(N), space complexity O(1)
func useFormular(coors [][]int) bool {
    n := len(coors)
    if n < 2 {
        return false
    }
    deltax, deltay := coors[1][0]-coors[0][0], coors[1][1]-coors[0][1]
    var ff f
    if deltax == 0 {
        ff = func(x int) int {
            return x
        }
    } else {
        slope := deltay / deltax
        b := coors[0][1] - slope*coors[0][0]
        ff = func(x int) int {
            return slope*x + b
        }
    }
    for i := 2; i < n; i++ {
        if deltax == 0 {
            if ff(coors[i][0]) != coors[i][0] {
                return false
            }
        } else {
            if ff(coors[i][0]) != coors[i][1] {
                return false
            }
        }
    }
    return true
}
