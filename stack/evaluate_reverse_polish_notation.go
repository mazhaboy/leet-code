package stack

import (
	"strconv"
)

func EvalRPN(tokens []string) int {
	operatorsMap := map[string]uint8{"*": '*', "+": '+', "-": '-', "/": '/'}
	res := 0
	stack := Constructor()
	for i := range tokens {
		operator, ok := operatorsMap[tokens[i]]
		if ok {
			num1 := stack.Top()
			stack.Pop()
			num2 := stack.Top()
			stack.Pop()
			if operator == '*' {
				res = num1 * num2
			} else if operator == '/' {
				res = num2 / num1
			} else if operator == '+' {
				res = num1 + num2
			} else if operator == '-' {
				res = num2 - num1
			}
			stack.Push(res)
		} else {
			num, _ := strconv.Atoi(tokens[i])
			stack.Push(num)
		}
	}
	res = stack.Top()

	return res
}
