package solutions

func RemoveElement(nums []int, val int) int {
	var i, j int
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
