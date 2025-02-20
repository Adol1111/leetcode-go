package solutions

func CombinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(path []int, index int, sum int)
	backtrack = func(path []int, index int, sum int) {
		if sum == target {
			// 需要深拷贝path，否则path会被修改
			res = append(res, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}

		// i从index开始，避免重复计算
		for i := index; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				continue
			}
			path = append(path, candidates[i])
			backtrack(path, i, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtrack(path, 0, 0)
	return res
}
