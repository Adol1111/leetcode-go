package solutions

func RemoveDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	count := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[count-2] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
