package solutions

func Calculate(s string) int {
	stack := make([]int, 0)
	sign := 1
	result := 0
	num := 0

	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = 1
			i++
		case '-':
			sign = -1
			i++
		case '(':
			stack = append(stack, result)
			stack = append(stack, sign)
			result = 0
			sign = 1
			i++
		case ')':
			result = stack[len(stack)-2] + stack[len(stack)-1]*result
			stack = stack[:len(stack)-2]
			i++
		default:
			num = 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			result += sign * num
		}
	}
	return result
}
