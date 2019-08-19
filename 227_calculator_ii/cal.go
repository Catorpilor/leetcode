package cal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/catorpilor/leetcode/utils"
)

func Calculate(input string) int {
	var ret int
	input = strings.Replace(input, " ", "", -1)
	st1, st2 := utils.NewStack(), utils.NewStack()
	bp, ep := len(input), len(input)
	ended := true
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fmt.Printf("current index: %d\n", i)
			if ended {
				bp = i
				ended = false
			}
			ep = i + 1
			if ep >= len(input) {
				st1.Push(input[bp:])
			}
		case '*', '/':
			ended = true
			var op1 string
			if ep >= len(input) {
				op1 = input[bp:]
			} else {
				op1 = input[bp:ep]
			}

			// st1.Push(input[bp:ep])
			if !st2.IsEmpty() {
				oprand2, _ := strconv.Atoi(op1)
				op := st2.Pop().(byte)
				oprand1, _ := strconv.Atoi(st1.Pop().(string))
				var ret int
				switch op {
				case '*':
					ret = oprand1 * oprand2
				case '/':
					ret = oprand1 / oprand2
				}
				st1.Push(strconv.Itoa(ret))
			} else {
				st1.Push(op1)
			}
			st2.Push(input[i])
		case '+', '-':
			ended = true
			var op1 string
			if ep >= len(input) {
				op1 = input[bp:]
			} else {
				op1 = input[bp:ep]
			}
			if !st2.IsEmpty() {
				oprand2, _ := strconv.Atoi(op1)
				oprand1, _ := strconv.Atoi(st1.Pop().(string))
				var res int
				op := st2.Pop().(byte)
				switch op {
				case '*':
					res = oprand1 * oprand2
				case '/':
					res = oprand1 / oprand2
				}
				st1.Push(strconv.Itoa(res))
			} else {
				st1.Push(op1)
			}
			if st1.Len() == 3 {
				op2, _ := strconv.Atoi(st1.Pop().(string))
				op := st1.Pop().(string)
				op1, _ := strconv.Atoi(st1.Pop().(string))
				switch op {
				case "+":
					st1.Push(strconv.Itoa(op1 + op2))
				case "-":
					st1.Push(strconv.Itoa(op1 - op2))
				}
			}
			// st1.Push(input[bp:ep])
			st1.Push(string(input[i]))
		}
	}
	fmt.Printf("stack1: %s\n", st1.String())
	fmt.Printf("stack2: %s\n", st2.String())
	if !st2.IsEmpty() {
		op2, _ := strconv.Atoi(st1.Pop().(string))
		op := st2.Pop().(byte)
		op1, _ := strconv.Atoi(st1.Pop().(string))
		switch op {
		case '*':
			ret = op1 * op2
		case '/':
			ret = op1 / op2
		}
		st1.Push(strconv.Itoa(ret))
	}
	for st1.Len() >= 3 {
		op2, _ := strconv.Atoi(st1.Pop().(string))
		op := st1.Pop().(string)
		op1, _ := strconv.Atoi(st1.Pop().(string))
		switch op {
		case "+":
			st1.Push(strconv.Itoa(op1 + op2))
		case "-":
			st1.Push(strconv.Itoa(op1 - op2))
		}
		fmt.Printf("st1 now: %s\n", st1.String())
	}
	if st1.Len() == 1 {
		ret, _ = strconv.Atoi(st1.Pop().(string))
	}

	return ret
}
