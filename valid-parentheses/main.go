package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	for _, c := range s {
		switch c {
		case '{':
			fmt.Println(string(c))
			stack = append(stack, c)
		case '(':
			fmt.Println(string(c))
			stack = append(stack, c)
		case '[':
			fmt.Println(string(c))
			stack = append(stack, c)
		case '}':
			fmt.Println("}", stack, len(stack))
			l := len(stack)
			if l == 0 || stack[l-1] != '{' {
				return false
			}
			stack = stack[:l-1]
			fmt.Println("} done", stack, len(stack))
		case ']':
			fmt.Println("]", stack, len(stack))
			l := len(stack)
			if l == 0 || stack[l-1] != '[' {
				return false
			}
			stack = stack[:l-1]
			fmt.Println("] done", stack, len(stack))
		case ')':
			fmt.Println(stack, len(stack))
			l := len(stack)
			if l == 0 || stack[l-1] != '(' {
				return false
			}
			stack = stack[:l-1]
			fmt.Println(stack, len(stack))
		default:
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isValid("{[]}"))
}
