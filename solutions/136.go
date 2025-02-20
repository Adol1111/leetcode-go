package solutions

func SingleNumber(nums []int) int {
	var result int
	for _, num := range nums {
		// 相同的数字异或为0，0与任何数字异或为任何数字本身
		// 根据题目只有一个数字出现一次，其他数字都出现两次
		// 所以把所有数字异或起来，相同的数字异或为0，0与任何数字异或为任何数字本身
		result ^= num
	}
	return result
}
