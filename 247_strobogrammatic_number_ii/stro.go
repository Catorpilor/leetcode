package stro

import "fmt"

func FindStro(n int) []string {
	return helper(n, n)
}

func FindStro2(n int) []string {
	var ans, cur []string
	if n&1 == 0 {
		ans = []string{""}
	} else {
		ans = []string{"0", "1", "8"}
	}
	if n < 2 {
		return ans
	}
	for ; n > 1; n -= 2 {
		cur = ans
		fmt.Println(cur)
		ans = nil
		for _, s := range cur {
			if n > 3 {
				ans = append(ans, "0"+s+"0")
			}
			ans = append(ans, "1"+s+"1")
			ans = append(ans, "6"+s+"9")
			ans = append(ans, "8"+s+"8")
			ans = append(ans, "9"+s+"6")
		}
	}
	return ans
}

func helper(curN, n int) []string {
	if curN == 0 {
		return []string{""}
	}
	if curN == 1 {
		return []string{"0", "1", "8"}
	}
	sub := helper(curN-2, n)
	var ret []string
	for _, c := range sub {
		if curN != n {
			ret = append(ret, "0"+c+"0")
		}
		ret = append(ret, "1"+c+"1")
		ret = append(ret, "6"+c+"9")
		ret = append(ret, "8"+c+"8")
		ret = append(ret, "9"+c+"6")
	}
	return ret

}
