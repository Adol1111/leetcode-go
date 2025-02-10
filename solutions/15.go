package solutions

import "sort"

func ThreeSum(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) < 3 {
		return ans
	}

	sort.Ints(nums)
	// 1. 变为有序数组
	// 2. 然后通过0-nums[i]=nums[j]+nums[k] 转化成两数之和的问题
	for i := 0; i < len(nums)-2; i++ {
		// 因为已经从小大大排列，所以如果当前元素大于0，那么后面的元素肯定都大于0，所以不可能存在三数之和为0
		if nums[i] > 0 {
			break
		}
		// 存在重复的元素，所以需要跳过
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
				// j去重
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				// k去重
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}

	return ans
}
