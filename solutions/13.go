package solutions

func RomanToInt(s string) int {
	sum := 0
	romanToIntMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romanToIntMap[s[i]] < romanToIntMap[s[i+1]] {
			sum -= romanToIntMap[s[i]]
		} else {
			sum += romanToIntMap[s[i]]
		}
	}

	return sum
}
