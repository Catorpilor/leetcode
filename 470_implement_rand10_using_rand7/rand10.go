package rand10

import "math/rand"

func rand10() int {
	return useFormular()
}

// useFormular time complexity O(1), space complexity O(1)
func useFormular() int {
	res := 40
	for res >= 40 {
		res = 7*(rand7()-1) + (rand7() - 1)
	}
	return res%10 + 1
}

func rand7() int {
	return rand.Intn(7) + 1
}
