package solutions

func PlusOne(digits []int) []int {
	// 数字是从高位到低位的，所以从后往前遍历
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 如果是9就要进位，所以把当前位设置为0，然后继续遍历
		digits[i] = 0
	}
	// 如果是 9999 这种，全部设置为0后，要在最前面补上1
	digits = append([]int{1}, digits...)
	return digits
}
