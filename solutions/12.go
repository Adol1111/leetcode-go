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
