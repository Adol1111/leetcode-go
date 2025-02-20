package solutions

func Permute(nums []int) [][]int {
	// 回溯
	var res [][]int
	var path []int
	used := make([]bool, len(nums))
	var backtrack func(path []int)
	backtrack = func(path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack(path)
	return res
}
