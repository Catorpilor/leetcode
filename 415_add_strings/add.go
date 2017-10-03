package add

import "fmt"

func AddStrings(num1, num2 string) string {
	n1, n2 := len(num1), len(num2)
	n := n1
	if n < n2 {
		n = n2
	}
	ret := make([]byte, 0, n+1)
	carrier := 0
	var x, y, sum, digit int
	for i, j := n1-1, n2-1; i >= 0 || j >= 0 || carrier == 1; i, j = i-1, j-1 {
		x, y = 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		sum = x + y + carrier
		digit = sum % 10
		ret = append(ret, byte(digit+'0'))
		carrier = sum / 10
	}
	fmt.Println(string(ret))
	return string(reverse(ret))
}

func reverse(sl []byte) []byte {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return sl
}
