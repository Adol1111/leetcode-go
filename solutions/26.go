package solutions

func RemoveDuplicates(nums []int) int {
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[count-1] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
