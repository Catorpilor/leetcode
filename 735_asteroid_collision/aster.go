package aster

func asterCollision(asteroids []int) []int {
	return useStack(asteroids)
}

// useStack time complexity O(N), space compelxity O(1)
func useStack(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	st := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nst := len(st)
		f := false
		for nst > 0 && st[nst-1]*arr[i] < 0 {
			if st[nst-1] > 0 {
				if st[nst-1]+arr[i] <= 0 {
					top := st[nst-1]
					nst--
					st = st[:nst]
					if top+arr[i] == 0 {
						// both explodes
						f = true
						break
					}
				} else {
					// arr[i] explode
					f = true
					break
				}
			} else {
				// top <-, arr[i] ->
				// never collision
				break
			}
		}
		if !f {
			st = append(st, arr[i])
		}

	}
	return st
}
