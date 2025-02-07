package solutions

import "sort"

func ThreeSum(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) < 3 {
		return ans
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
				continue
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
				continue
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}

	return ans
}
