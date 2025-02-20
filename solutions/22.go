package solutions

func GenerateParenthesis(n int) []string {
	var res []string
	var backtrack func(path string, left int, right int)
	backtrack = func(path string, left int, right int) {
		if left == n && right == n {
			res = append(res, path)
			return
		}
		if left < n {
			backtrack(path+"(", left+1, right)
		}
		// 右括号数量不能超过左括号，不然就会出现 )( 的情况
		if right < left {
			backtrack(path+")", left, right+1)
		}
	}
	backtrack("", 0, 0)
	return res
}
