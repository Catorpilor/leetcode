package cal

import (
	"fmt"
	"math"
	"strings"

	"github.com/catorpilor/LeetCode/utils"
)

// Calculator calculate s
func Calculator(s string) int {
	st := utils.NewStack()
	return reverseWithStack(s, st)
}

func reverseWithStack(s string, st *utils.Stack) int {
	s = strings.Replace(s, " ", "", -1)
	var n, op int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			op = int(math.Pow10(n))*(int(s[i]-'0')) + op
			n++
		} else {
			if n != 0 {
				st.Push(op)
				op = 0
				n = 0
			}
			switch s[i] {
			case '(':
				res := evalutateSt(st)
				st.Pop()
				st.Push(res)
			default:
				st.Push(s[i])
			}
		}
	}
	if n != 0 {
		st.Push(op)
	}
	return evalutateSt(st)
}

func evalutateSt(st *utils.Stack) int {
	fmt.Printf("before evaluate stack is %s\n", st.String())
	var res int
	if !st.IsEmpty() {
		res = st.Pop().(int)
	}
	for !st.IsEmpty() {
		sign := st.Pop().(byte)
		if sign == '+' {
			res += st.Pop().(int)
		} else if sign == '-' {
			res -= st.Pop().(int)
		} else {
			// it's ')'
			// st.Pop()
			break
		}
	}
	fmt.Printf("after evaluate stack is %s\n", st.String())

	return res
}
