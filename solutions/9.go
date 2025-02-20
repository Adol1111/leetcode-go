package solutions

func IsPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	origin := x
	var reverse int
	// 从个位到最高位，反过来重新组合，如果和原数相同则是回文
	for x != 0 {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return origin == reverse
}
