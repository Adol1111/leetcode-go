package solutions

func TwoSum(numbers []int, target int) []int {
	ans := make([]int, 2)
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
