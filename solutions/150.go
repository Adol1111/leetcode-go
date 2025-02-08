package solutions

import "strconv"

func EvalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	var stack []int
	operator := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	for _, token := range tokens {
		if operator[token] {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			var result int
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}
			stack = append(stack, result)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
