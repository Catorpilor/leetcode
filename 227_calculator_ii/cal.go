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
	st1 := utils.NewStack()
	var highOp string
	bp, ep := len(input), len(input)
	ended := true
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if ended {
				bp = i
				ended = false
			}
			ep = i + 1
			if ep >= len(input) {
				st1.Push(input[bp:])
			}
		case '*', '/', '+', '-':
			ended = true
			var op2 string
			if ep >= len(input) {
				op2 = input[bp:]
			} else {
				op2 = input[bp:ep]
			}
			var higherOp bool
			ops := string(input[i])
			if ops == "*" || ops == "/" {
				higherOp = true
			}
			highOp = cal(st1, ops, highOp, op2, higherOp)
			// st1.Push(input[bp:ep])

			if higherOp {
				highOp = ops
			} else {
				st1.Push(ops)
			}
		}
	}
	fmt.Printf("stack1: %s\n", st1.String())
	// fmt.Printf("stack2: %s\n", st2.String())
	if highOp != "" {
		op2, _ := strconv.Atoi(st1.Pop().(string))
		// op := st2.Pop().(string)
		op1, _ := strconv.Atoi(st1.Pop().(string))
		switch highOp {
		case "*":
			ret = op1 * op2
		case "/":
			ret = op1 / op2
		}
		highOp = ""
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

func cal(st1 *utils.Stack, op, prevOp, op2 string, higherOp bool) string {
	fmt.Printf("stack1 is %s, op is %s, prevOp is %s, op2 is %s and flag is %t\n",
		st1.String(), op, prevOp, op2, higherOp)
	// if !st2.IsEmpty() {
	if prevOp != "" {
		oprand2, _ := strconv.Atoi(op2)
		oprand1, _ := strconv.Atoi(st1.Pop().(string))
		var res int
		// op := st2.Pop().(string)
		switch prevOp {
		case "*":
			res = oprand1 * oprand2
		case "/":
			res = oprand1 / oprand2
		}
		prevOp = ""
		st1.Push(strconv.Itoa(res))
	} else {
		st1.Push(op2)
	}
	if !higherOp {
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
	}
	return prevOp
}
