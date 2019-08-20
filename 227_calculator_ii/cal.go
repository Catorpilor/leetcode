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

/*
curInde: 0, num: 1, lastSign: +, sum: 0, tempSum: 0
curInde: 1, num: 1, lastSign: +, sum: 0, tempSum: 0
curInde: 2, num: 2, lastSign: *, sum: 0, tempSum: 1
curInde: 3, num: 2, lastSign: *, sum: 0, tempSum: 1
curInde: 4, num: 3, lastSign: -, sum: 0, tempSum: 2
curInde: 5, num: 3, lastSign: -, sum: 0, tempSum: 2
curInde: 6, num: 4, lastSign: /, sum: 2, tempSum: -3
curInde: 7, num: 4, lastSign: /, sum: 2, tempSum: -3
curInde: 8, num: 5, lastSign: +, sum: 2, tempSum: 0
curInde: 9, num: 5, lastSign: +, sum: 2, tempSum: 0
curInde: 10, num: 6, lastSign: *, sum: 2, tempSum: 5
curInde: 11, num: 6, lastSign: *, sum: 2, tempSum: 5
curInde: 12, num: 7, lastSign: -, sum: 2, tempSum: 30
curInde: 13, num: 7, lastSign: -, sum: 2, tempSum: 30
curInde: 14, num: 8, lastSign: *, sum: 32, tempSum: -7
curInde: 15, num: 8, lastSign: *, sum: 32, tempSum: -7
curInde: 16, num: 9, lastSign: +, sum: 32, tempSum: -56
curInde: 17, num: 9, lastSign: +, sum: 32, tempSum: -56
curInde: 18, num: 1, lastSign: /, sum: -24, tempSum: 9
curInde: 19, num: 10, lastSign: /, sum: -24, tempSum: 9
*/

func Calculate2(s string) int {
	s = strings.Replace(s, " ", "", -1)
	var sum, tempSum, num int
	lastSign := byte('+')
	for i := 0; i < len(s); i++ {
		// fmt.Printf("curInde: %d, num: %d, ")
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		}
		if i == len(s)-1 || s[i] < '0' || s[i] > '9' {
			switch lastSign {
			case '+':
				sum += tempSum
				tempSum = num
			case '-':
				sum += tempSum
				tempSum = -num
			case '*':
				tempSum *= num
			case '/':
				tempSum /= num
			}
			lastSign = s[i]
			num = 0
		}
		fmt.Printf("curInde: %d, num: %d, lastSign: %s, sum: %d, tempSum: %d\n",
			i, num, string(lastSign), sum, tempSum)
	}
	sum += tempSum
	return sum
}
