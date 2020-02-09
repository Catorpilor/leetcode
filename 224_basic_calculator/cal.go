package cal

import (
	"fmt"
	"math"
	"strings"

	"github.com/catorpilor/leetcode/utils"
)

// Calculator calculate s
func Calculator(s string) int {
	st := utils.NewStack()
	// return reverseWithStack(s, st)
	return iterativeWithStack(s, st)
}

// based on discussion:
// https://leetcode.com/problems/basic-calculator/discuss/62361/southpenguin
func iterativeWithStack(s string, st *utils.Stack) int {
	s = strings.Replace(s, " ", "", -1)
	sign := 1
	operand := 0
	var result int
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			operand = operand*10 + int(s[i]-'0')
		} else if s[i] == '+' {
			result += sign * operand
			operand = 0
			sign = 1
		} else if s[i] == '-' {
			result += sign * operand
			operand = 0
			sign = -1
		} else if s[i] == '(' {
			st.Push(result)
			st.Push(sign)
			// reset sing and oprand
			sign = 1
			result = 0
		} else {
			// ')'
			result += sign * operand
			operand = 0
			result *= st.Pop().(int) // *sign
			result += st.Pop().(int) // the pre calculated result before '('
		}
	}
	return result + sign*operand
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
				// st.Pop()
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
