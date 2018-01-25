package find

import (
	"math"
	"strings"
)

func Match(a, b string) int {
	la, lb := len(a), len(b)
	i, j := 0, 0
	for i < la {
		j = 0
		for j < lb && a[(i+j)%la] == b[j] {
			j += 1
		}
		if j == lb {
			add := 0
			if (i+j)%la != 0 {
				add = 1
			}
			return (i+j)/la + add
		}
		i += 1
	}
	return -1
}

func Match2(A, B string) int {
	var repeats int

	for bi := 0; bi < len(B); {
		for ai := 0; ai < len(A); ai++ {
			if bi == len(B) {
				break
			}

			if B[bi] != A[ai] {
				if repeats > 0 {
					return -1
				}

				bi = 0
			} else {
				bi++
			}
		}

		repeats++
	}

	return repeats
}

func Match3(A, B string) int {
	// kmp create suffix and prefix index
	pf := make([]int, len(B))
	for i, j := 1, 0; i < len(B); i++ {
	KMP:
		if B[j] == B[i] {
			j += 1
			pf[i] = j
		} else {
			if j > 0 {
				j = pf[j-1]
				goto KMP
			}
		}
	}
	for i, j := 0, 0; i < len(A); i++ {
		if j > 0 {
			j = pf[j-1]
		}
		for j < len(B) && A[(i+j)%len(A)] == B[j] {
			j += 1
		}
		if j == len(B) {
			add := 0
			if (i+j)%len(A) != 0 {
				add = 1
			}
			return (i+j)/len(A) + add
		}
	}
	return -1
}

func Match4(a, b string) int {
	if len(a) == 0 {
		return -1
	}
	min := int(math.Max(float64(len(b)/len(a)), 1.0))
	if strings.Contains(strings.Repeat(a, min), b) {
		return min
	}
	if strings.Contains(strings.Repeat(a, min+1), b) {
		return min + 1
	}
	return -1
}

func Match5(a, b string) int {
	if len(a) == 0 {
		return -1
	}
	res := a
	for ret := 1; ret <= len(b)/len(a)+2; ret++ {
		if strings.Index(res, b) != -1 {
			return ret
		}
		res += a
	}
	return -1
}
