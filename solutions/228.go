package solutions

import "strconv"

func SummaryRanges(nums []int) []string {
	var res []string
	for i := 0; i < len(nums); {
		j := i + 1
		for j < len(nums) && nums[j] == nums[j-1]+1 {
			j++
		}
		if i == j-1 {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			res = append(res, strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[j-1]))
		}
	}
	return res
}
