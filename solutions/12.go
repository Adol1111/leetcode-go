package solutions

func IntToRoman(num int) string {
	var valueSymbols = []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	// 先处理大的数，再处理小的数
	// 循环一直减，直到num小于value，然后继续下一个value
	// 如果减到0，则直接返回
	ans := ""
	for _, v := range valueSymbols {
		for num >= v.value {
			num -= v.value
			ans += v.symbol
		}
		if num == 0 {
			break
		}
	}
	return ans
}
