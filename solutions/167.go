package solutions

func TwoSum(numbers []int, target int) []int {
	ans := make([]int, 2)
	// 因为是有序的，所以可以使用双指针，一个指向头，一个指向尾
	// 左侧为最小值，右侧为最大值，如果两个值相加小于target，则左侧指针右移，如果大于target，则右侧指针左移
	// 如果i>=j说明已经匹配完，不需要再继续
	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] == target {
			ans[0] = i + 1
			ans[1] = j + 1
			break
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return ans
}
