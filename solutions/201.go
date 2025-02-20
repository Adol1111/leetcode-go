package solutions

func rangeBitwiseAnd(left int, right int) int {
	// 101
	// 110
	// 111

	// 100

	shift := 0
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}
