package solutions

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	// 定义数字对应的字母
	digitMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	// 定义结果
	var res []string

	// 定义回溯函数
	var backtrack func(digits string, index int, path string)
	backtrack = func(digits string, index int, path string) {
		// 如果已经遍历完了，就把路径加入结果
		if index == len(digits) {
			res = append(res, path)
			return
		}
		// 获取当前数字对应的字母
		letters := digitMap[digits[index]]
		// 遍历字母
		for i := 0; i < len(letters); i++ {
			// 递归
			backtrack(digits, index+1, path+string(letters[i]))
		}
	}
	// 调用回溯函数
	backtrack(digits, 0, "")
	return res
}
