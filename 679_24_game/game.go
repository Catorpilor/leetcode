package game

import (
	"math"
)

const eps = 0.001

func JudgePoint(nums []int) bool {
	// convert to float64
	res := make([]float64, 0, len(nums))
	for i := range nums {
		res = append(res, float64(nums[i]))
	}
	return helper(res)

}

func helper(res []float64) bool {
	if len(res) == 0 {
		return false
	}
	if len(res) == 1 {
		return math.Abs(res[0]-24.0) < eps
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < i; j++ {
			// p1+p2, p1-p2, p2-p1, p1*p2, p1/p2, p2/p1
			next := make([]float64, 0, 6)
			pi, pj := res[i], res[j]
			next = append(next, pi+pj, pi-pj, pj-pi, pi*pj)
			if math.Abs(pi) > eps {
				next = append(next, pj/pi)
			}
			if math.Abs(pj) > eps {
				next = append(next, pi/pj)
			}
			copy(res[i:], res[i+1:])
			res[len(res)-1] = 0.0 // or the zero value of T
			res = res[:len(res)-1]
			copy(res[j:], res[j+1:])
			res[len(res)-1] = 0.0 // or the zero value of T
			res = res[:len(res)-1]

			for _, v := range next {
				n := len(res)
				res = append(res, v)
				helper(res)
				res = resFCGFDERCX FDDDDD          S

			}

		}
	}
	return false
}
